package compute

import "time"

func (*Compute) MakeTodayStart() time.Time {
	tmp := time.Now()
	return time.Date(tmp.Year(), tmp.Month(), tmp.Day(), 0, 0, 0, 0, time.Local)
}
func (*Compute) MakeTodayEnd() time.Time {
	tmp := time.Now()
	return time.Date(tmp.Year(), tmp.Month(), tmp.Day(), 23, 59, 59, 0, time.Local)
}
