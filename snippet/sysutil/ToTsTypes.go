package structutil

import (
	"fmt"
	"reflect"
)

// ToTsTypes 结构体转换为ts类型
// 未指定 y-ts-types-spacer 时 默认间隔符号 :
// 指定 y-ts-types 为 - 时不转换 不为 - json : y-ts-types
// 未指定 y-ts-types 时 json : goTypeToTsTypes(v.Field(k).Type().String())
// 未指定 y-ts-types 且 json 为 -或为空 时不转换
func ToTsTypes(obj interface{}, isModel bool) (filename, row string) {
	t := reflect.TypeOf(obj)
	filename = fmt.Sprintf("%v.d.ts", t.Name())
	v := reflect.ValueOf(obj)
	row = fmt.Sprintf("interface %v {\n", t.Name())
	var appendOtherExtends []string
	for k := 0; k < t.NumField(); k++ {
		yRLabel := t.Field(k).Tag.Get("json")
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
		row += fmt.Sprintf("interface %v extends %v {}", t.Name(), extend)
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
