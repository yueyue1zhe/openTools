package compute

import (
	"fmt"
	"github.com/shopspring/decimal"
	"strconv"
)

type Compute struct {
}

func NewCompute() *Compute {
	return &Compute{}
}

// Decimal2Init 浮点数保留2位小数点4舍5入
func (y *Compute) Decimal2Init(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

// Percent 整数转换为百分比的小数
func (y *Compute) Percent(val int64) float64 {
	a, _ := decimal.NewFromInt(val).Float64()
	b, _ := decimal.NewFromInt(100).Float64()
	return y.NewPriceCompute(a, b).Div()
}

func (y *Compute) NewPriceCompute(a, b float64) *priceCompute {
	fl1 := decimal.NewFromFloat(a)
	fl2 := decimal.NewFromFloat(b)
	return &priceCompute{a: fl1, b: fl2}
}
