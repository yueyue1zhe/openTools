package decimalutil

import (
	"e.coding.net/zhechat/magic/taihao/function/commutil"
	"fmt"
	"github.com/shopspring/decimal"
	"strconv"
)

// FloatDecimal2Init Decimal2Init 浮点数保留2位小数点4舍5入
func FloatDecimal2Init(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

// FloatPercent Percent 整数转换为百分比的小数
func FloatPercent(val int64) float64 {
	a, _ := decimal.NewFromInt(val).Float64()
	b, _ := decimal.NewFromInt(100).Float64()
	return NewPriceCompute(a, b).Div()
}

func RandomFloat64(min, max float64) (out float64) {
	if min == max {
		return min
	}
	dMin := NewPriceCompute(min, 100).Mul()
	dMax := NewPriceCompute(max, 100).Mul()
	oRaw := commutil.Int64Random(int64(dMin), int64(dMax))
	o := NewPriceCompute(float64(oRaw), 100).Div()
	out = FloatDecimal2Init(o)
	if out > max || out < min || out == 0 {
		return RandomFloat64(min, max)
	}
	return out
}
