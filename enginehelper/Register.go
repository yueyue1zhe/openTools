package aiCreator

import (
	"e.coding.net/zhechat/magic/taihao/library/dbutil"
)

func Register() error {
	if err := install(); err != nil {
		return err
	}
	upgrade()
	return nil
}

func tableTypes() []interface{} {
	return []interface{}{}
}
func OtherTypes() (list []interface{}) {
	list = append(list, tableTypes()...)
	return append(list, []interface{}{}...)
}
func FrontendOtherTypes() []interface{} {
	return []interface{}{}
}
func install() error {
	for _, i := range tableTypes() {
		if err := dbutil.ModelCreateIfNotExistWithErr(i); err != nil {
			return err
		}
	}
	return nil
}

func upgrade() {

}
