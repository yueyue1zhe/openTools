package compute

import (
	"bytes"
	"crypto/rand"
	"math/big"
)

type Random struct {
}

func (*Compute) NewRandom() *Random {
	return &Random{}
}

func (*Random) String(len int) string {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyABCDEFGHIJKLMNOPQRSTUVWXY"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}
func (y *Random) Int64(min, max int64) int64 {
	maxBigInt := big.NewInt(max)
	i, _ := rand.Int(rand.Reader, maxBigInt)
	if i.Int64() < min {
		y.Int64(min, max)
	}
	return i.Int64()
}

// Float64 生成指定范围内随机浮点数 返回结果4舍5入保留小数点后2位
func (y *Random) Float64(min, max float64) (out float64) {
	if min == max {
		return min
	}
	compute := NewCompute()
	dMin := compute.NewPriceCompute(min, 100).Mul()
	dMax := compute.NewPriceCompute(max, 100).Mul()
	oRaw := y.Int64(int64(dMin), int64(dMax))
	o := compute.NewPriceCompute(float64(oRaw), 100).Div()
	out = compute.Decimal2Init(o)
	if out > max || out < min || out == 0 {
		return y.Float64(min, max)
	}
	return out
}
