package commutil

import "reflect"

// SliceNormal 工具函数处理的简单切片类型
type SliceNormal interface {
	int | int64 | string | uint
}

// SliceIncludes 判断切片是否包含指定元素
func SliceIncludes[T SliceNormal](target T, sucList []T) bool {
	for _, t := range sucList {
		if target == t {
			return true
		}
	}
	return false
}

// SliceRemoveRepeated 删除切片中重复元素
func SliceRemoveRepeated[T SliceNormal](s []T) []T {
	var result []T
	m := make(map[T]bool) //map的值不重要
	for _, v := range s {
		if _, ok := m[v]; !ok {
			result = append(result, v)
			m[v] = true
		}
	}
	return result
}

// SliceRemoveTarget 删除切片中目标元素
func SliceRemoveTarget[T SliceNormal](target T, list []T) []T {
	index := -1
	for i, s := range list {
		if target == s {
			index = i
		}
	}
	if index == -1 {
		return list
	}
	return SliceRemoveIndex(index, list)
}

// SliceRemoveIndex 删除切片中指定下标元素
func SliceRemoveIndex[T SliceNormal](index int, list []T) []T {
	return append(list[:index], list[index+1:]...)
}

// SliceToIds 提取切片中字段名称为ID的值 组成新的切片
func SliceToIds(list interface{}) []uint {
	return SliceToUintArr(list, "ID")
}

// SliceToUintArr 提取切片中字段名称为ID的值 组成新的切片
func SliceToUintArr(list interface{}, field string) (out []uint) {
	for i := 0; i < reflect.ValueOf(list).Len(); i++ {
		out = append(out, uint(reflect.ValueOf(list).Index(i).FieldByName(field).Uint()))
	}
	return
}
