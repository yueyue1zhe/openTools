package defined

import (
	"e.coding.net/zhechat/magic/taihao/function/commutil"
	"fmt"
)

type Gender int

const (
	GenderUnknown Gender = iota
	GenderMan
	GenderWoman
)

func (g Gender) CustomInt() int {
	return int(g)
}
func (g Gender) Check() error {
	suc := []Gender{GenderMan, GenderWoman}
	if !commutil.CustomIntSliceIncludes(g, suc) {
		return fmt.Errorf("请正确设置性别")
	}
	return nil
}

type MobilePhone string

func (m MobilePhone) CustomString() string {
	return string(m)
}
func (m MobilePhone) MaskShow() string {
	if m.CustomString() == "" {
		return ""
	}
	pre := commutil.StringCut(m.CustomString(), 0, 3)
	end := commutil.StringCut(m.CustomString(), 7, 11)
	return pre + "****" + end
}
