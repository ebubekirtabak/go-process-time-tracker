package types

import "time"

type Measure struct {
	Name string
	StartTime int64
	FinishTime int64
	PauseStartTime int64
	TotalTime time.Duration
	Delaylogs []DelayLog
}

type DelayLog struct {
	StartTime int64
	FinishTime int64
}