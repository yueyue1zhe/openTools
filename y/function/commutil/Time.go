package commutil

import (
	"github.com/shopspring/decimal"
	"time"
)

func TimeTodayStart() time.Time {
	tmp := time.Now()
	return time.Date(tmp.Year(), tmp.Month(), tmp.Day(), 0, 0, 0, 0, time.Local)
}
func TimeTodayEnd() time.Time {
	tmp := time.Now()
	return time.Date(tmp.Year(), tmp.Month(), tmp.Day(), 23, 59, 59, 0, time.Local)
}

func SecondToDay(s int) float64 {
	useS := float64(s)
	if s == 0 {
		return useS
	}
	fl1 := decimal.NewFromFloat(useS)
	fl2 := decimal.NewFromFloat(float64(60 * 60 * 24))
	return fl1.Div(fl2).RoundFloor(2).InexactFloat64()
}
