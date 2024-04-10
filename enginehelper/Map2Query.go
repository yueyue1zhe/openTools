package compute

import (
	"fmt"
	"net/url"
)

// Map2Query map生成query a=b&c=d
func Map2Query(param map[string]interface{}) string {
	query := url.Values{}
	for s, i := range param {
		query.Add(s, fmt.Sprintf("%v", i))
	}
	return query.Encode()
}
