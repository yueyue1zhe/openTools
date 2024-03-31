package encrypt

import (
	"crypto/md5"
	"fmt"
	"sort"
	"strings"
)

func SignMake(data map[string]any, secret string) string {
	var dataParams []string
	var keys []string
	for k := range data {
		if k != "sign" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	for _, key := range keys {
		dataParams = append(dataParams, fmt.Sprintf("%v=%v", key, data[key]))
	}
	srcCode := md5.Sum([]byte(strings.Join(dataParams, "&") + secret))
	return fmt.Sprintf("%x", srcCode)
}
