package zfx

import (
	"time"
)

type _absoluteTime struct {
}

func (inst *_absoluteTime) CurrDate() string {
	return time.Now().Format("2006-01-02")
}

func (inst *_absoluteTime) CurrTime() string {
	return time.Now().Format("15:04:05")
}

func (inst *_absoluteTime) CurrUtc() int {
	return int(time.Now().Unix())
}

func (inst *_absoluteTime) CurrUtcTo(offsetSeconds int) int {
	return int(time.Now().Add(time.Duration(offsetSeconds) * time.Second).Unix())
}

func (inst *_absoluteTime) Sleep(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}
