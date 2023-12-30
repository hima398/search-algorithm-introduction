package util

import (
	"time"
)

type TimeKeeper struct {
	StartTime     int64
	TimeThreshold int64
}

func NewTimeKeeper(milliSecond int) *TimeKeeper {
	res := new(TimeKeeper)
	res.StartTime = time.Now().UnixMicro()
	res.TimeThreshold = int64(1000 * milliSecond)
	return res
}

func (t *TimeKeeper) IsTimeOver() bool {
	diff := time.Now().UnixMicro() - t.StartTime
	return diff >= t.TimeThreshold
}
