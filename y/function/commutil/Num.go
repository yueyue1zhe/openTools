package commutil

import (
	"crypto/rand"
	"math"
	"math/big"
)

// Int64Random 生成指定范围的随机数
func Int64Random(min, max int64) int64 {
	maxBigInt := big.NewInt(max)
	i, _ := rand.Int(rand.Reader, maxBigInt)
	if i.Int64() < min {
		Int64Random(min, max)
	}
	return i.Int64()
}

type TaiHaoNum interface {
	int64 | int
}

func NumMax[T TaiHaoNum](o, t T) T {
	return T(math.Max(float64(o), float64(t)))
}

func NumDefault[T TaiHaoNum](o, d T) T {
	if o > 0 {
		return o
	}
	return d
}
