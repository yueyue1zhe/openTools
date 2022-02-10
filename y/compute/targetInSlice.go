package compute

func (*Compute) StringInSlice(target string, sucList []string) bool {
	for _, s := range sucList {
		if target == s {
			return true
		}
	}
	return false
}

func (*Compute) Int64InSlice(target int64, sucList []int64) bool {
	for _, s := range sucList {
		if target == s {
			return true
		}
	}
	return false
}

func (*Compute) UintInSlice(target uint, sucList []uint) bool {
	for _, s := range sucList {
		if target == s {
			return true
		}
	}
	return false
}
