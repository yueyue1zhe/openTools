package commutil

import (
	"e.coding.net/zhechat/magic/taihao/core"
	"fmt"
	"reflect"
	"strings"
)

// StructToMap 根据json tag 将struct 转map
func StructToMap(raw any) map[string]any {
	data := make(map[string]any)
	objT := reflect.TypeOf(raw)
	objV := reflect.ValueOf(raw)
	for i := 0; i < objT.NumField(); i++ {
		if name, ok := objT.Field(i).Tag.Lookup("json"); ok {
			data[name] = objV.Field(i).Interface()
		}
	}
	return data
}

func StructArrRequiredJudge(objs ...any) error {
	for _, obj := range objs {
		if err := StructRequiredJudge(obj); err != nil {
			return err
		}
	}
	return nil
}

// StructRequiredJudge 结构体 检测 yrl 标记字段 是否为空
func StructRequiredJudge(obj interface{}) error {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	appConf := core.GetAppRegisterConf()
	for k := 0; k < t.NumField(); k++ {
		if appConf.StructBasic != "" && t.Field(k).Tag.Get("json") == appConf.StructBasic && v.Field(k).IsZero() {
			return fmt.Errorf("基础平台参数异常")
		}
		yRLabel := t.Field(k).Tag.Get("y-required-label")
		if yRLabel != "" && v.Field(k).IsZero() {
			pre := t.Field(k).Tag.Get("y-required-pre")
			if pre == "" {
				return fmt.Errorf("请设置%v", yRLabel)
			}
			return fmt.Errorf("%v%v", pre, yRLabel)
		}
	}
	return nil
}

// StructToTagLabelSlice 结构体获取 指定类型 字段名称切片 [json | 其它tag标记]
func StructToTagLabelSlice(obj interface{}, tag string) []string {
	t := reflect.TypeOf(obj)
	var useTag []string
	for k := 0; k < t.NumField(); k++ {
		mapstructureTag := t.Field(k).Tag.Get(tag)
		if mapstructureTag != "" {
			useTag = append(useTag, mapstructureTag)
		}
	}
	return useTag
}

// StructToTsTypes 结构体转换为ts类型
// typePreName 指定时 结构体首字母大写 追加前缀
// 未指定 y-ts-types-spacer 时 默认间隔符号 :
// 指定 y-ts-types 为 - 时不转换 不为 - json : y-ts-types
// 未指定 y-ts-types 时 json : goTypeToTsTypes(v.Field(k).Type().String())
// 未指定 y-ts-types 且 json 为 -或为空 时不转换
// 20221211 增加针对sql.NullTime的类型处理
func StructToTsTypes(obj interface{}, typePreName string) (filename, row string) {
	t := reflect.TypeOf(obj)
	filename = fmt.Sprintf("%v.d.ts", t.Name())
	v := reflect.ValueOf(obj)
	useTypeName := t.Name()
	if typePreName != "" {
		useTypeNameFirstWorld := useTypeName[0:1]
		useTypeName = typePreName + strings.Replace(useTypeName, useTypeNameFirstWorld, strings.ToUpper(useTypeNameFirstWorld), 1)
	}
	row = fmt.Sprintf("interface %v {\n", useTypeName)
	var appendOtherExtends []string
	for k := 0; k < t.NumField(); k++ {
		yRLabel := t.Field(k).Tag.Get("json")

		if t.Name() == "NullTime" {
			suc := []string{"Time", "Valid"}
			if SliceIncludes(t.Field(k).Name, suc) {
				yRLabel = t.Field(k).Name
			}
		}

		yRType := v.Field(k).Type().String()
		yTsSpacer := t.Field(k).Tag.Get("y-ts-types-spacer")
		if yTsSpacer == "" {
			yTsSpacer = ":"
		}
		yTsTypes := t.Field(k).Tag.Get("y-ts-types")
		if yTsTypes != "" {
			if yTsTypes != "-" {
				row += fmt.Sprintf("    %v%v %v;\n", yRLabel, yTsSpacer, yTsTypes)
			}
			continue
		}
		tsType := goTypeToTsTypes(yRType)
		if yRLabel != "" && yRLabel != "-" && tsType != "" {
			row += fmt.Sprintf("    %v%v %v;\n", yRLabel, yTsSpacer, tsType)
			continue
		}
		if yRLabel != "-" {
			otherType := v.Field(k).Type().Name()
			if yRLabel == "" {
				appendOtherExtends = append(appendOtherExtends, otherType)
			} else {
				row += fmt.Sprintf("    %v%v %v;\n", yRLabel, yTsSpacer, otherType)
			}
		}
	}
	row += "}\n"
	for _, extend := range appendOtherExtends {
		row += fmt.Sprintf("interface %v extends %v {}\n", useTypeName, extend)
	}
	return
}

func goTypeToTsTypes(target string) string {
	m := make(map[string]string)
	m["float64"] = "number"
	m["int64"] = "number"
	m["int"] = "number"
	m["uint"] = "number"
	m["bool"] = "boolean"
	m["string"] = "string"
	m["time.Time"] = "string"
	return m[target]
}
