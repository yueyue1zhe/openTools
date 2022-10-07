package structutil

import (
	"fmt"
	"reflect"
)

func ToTsTypes(obj interface{}, isModel bool) (filename, row string) {
	t := reflect.TypeOf(obj)
	filename = fmt.Sprintf("%v.d.ts", t.Name())
	v := reflect.ValueOf(obj)
	row = fmt.Sprintf("interface %v {\n", t.Name())
	if isModel {
		row += "    id: number;\n"
		row += "    created_at: string;\n"
		row += "    updated_at: string;\n"
	}
	for k := 0; k < t.NumField(); k++ {
		yRLabel := t.Field(k).Tag.Get("json")
		yRType := v.Field(k).Type().String()
		goTYpe := goTypeToTsTypes(yRType)
		if yRLabel != "" && yRLabel != "-" && goTYpe != "" {
			row += fmt.Sprintf("    %v: %v;\n", yRLabel, goTYpe)
		}
	}
	row += "}"
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
