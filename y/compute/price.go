package compute

import (
	"github.com/shopspring/decimal"
)

type priceCompute struct {
	a decimal.Decimal
	b decimal.Decimal
}

func (m *priceCompute) Div() float64 {
	res, _ := m.a.Div(m.b).Float64()
	return NewCompute().Decimal2Init(res)
}

func (m *priceCompute) Mul() float64 {
	res, _ := m.a.Mul(m.b).Float64()
	return NewCompute().Decimal2Init(res)
}

func (m *priceCompute) Sub() float64 {
	res, _ := m.a.Sub(m.b).Float64()
	return NewCompute().Decimal2Init(res)
}

func (m *priceCompute) Sum() float64 {
	res, _ := m.a.Add(m.b).Float64()
	return NewCompute().Decimal2Init(res)
}
