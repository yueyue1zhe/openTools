package upgrade

import (
	"fmt"
	"reflect"
)

func upgradeAddColumn(mod interface{}, field string, insertFunc func(mod interface{})) {
	if !mysql.DB.Migrator().HasTable(&mod) {
		fmt.Println(fmt.Sprintf("table %v is not find", reflect.TypeOf(mod)))
		return
	}
	if !mysql.DB.Migrator().HasColumn(&mod, field) {
		if err := mysql.DB.Migrator().AddColumn(&mod, field); err != nil {
			fmt.Println(fmt.Sprintf("table %v add column is %v fail:", reflect.TypeOf(mod), field), err.Error())
		} else {
			insertFunc(mod)
		}
	}
}
