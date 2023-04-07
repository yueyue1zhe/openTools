package commutil

import (
	"fmt"
	"net/url"
)

// MapToQuery ToQuery map生成query a=b&c=d
func MapToQuery(param map[string]interface{}) string {
	query := url.Values{}
	for s, i := range param {
		query.Add(s, fmt.Sprintf("%v", i))
	}
	return query.Encode()
}

func MapMerge[T any](input ...map[string]T) map[string]T {
	out := make(map[string]T)
	for _, m := range input {
		for s, a := range m {
			out[s] = a
		}
	}
	return out
}
