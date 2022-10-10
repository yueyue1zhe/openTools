package stringutil

import (
	"strconv"
	"strings"
)

func SplitUint(str string) (res []uint) {
	raw := strings.Split(str, ",")
	if len(raw) > 0 {
		for _, s := range raw {
			if uR, err := strconv.ParseUint(s, 10, 64); err == nil {
				res = append(res, uint(uR))
			}
		}
	}
	return
}
