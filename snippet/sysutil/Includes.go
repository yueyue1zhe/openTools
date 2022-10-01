package sliceutil

func Includes[T int | int64 | string | uint](target T, sucList []T) bool {
	for _, t := range sucList {
		if target == t {
			return true
		}
	}
	return false
}
