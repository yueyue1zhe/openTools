package color

import (
	"image/color"
	"strconv"
	"strings"
)

func Hex2rgb(str string) color.RGBA {
	if len(str) < 7 || strings.Index(str, "#") != 0 {
		return color.RGBA{}
	}
	str = strings.TrimLeft(str, "#")
	r, _ := strconv.ParseUint(str[:2], 16, 10)
	g, _ := strconv.ParseInt(str[2:4], 16, 18)
	b, _ := strconv.ParseInt(str[4:], 16, 10)
	return color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: 255,
	}
}
