package maputil

import (
	"fmt"
	"net/url"
)

// ToQuery map生成query a=b&c=d
func ToQuery(param map[string]interface{}) string {
	query := url.Values{}
	for s, i := range param {
		query.Add(s, fmt.Sprintf("%v", i))
	}
	return query.Encode()
}
