package commutil

import (
	"fmt"
	"testing"
)

func TestMapToQuery(t *testing.T) {
	type common struct {
		A string `json:"a" y-required-label:"a"`
	}
	type use struct {
		common
		DDD string `json:"ddd"`
	}
	fmt.Println(StructRequiredJudge(use{}))
}

func TestTokenKey(t *testing.T) {
	fmt.Println(StringRandom(32))
	fmt.Println(StringRandom(43))
}
