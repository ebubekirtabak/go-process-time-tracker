package types

import "time"

type Measure struct {
	Name string
	StartTime int64
	FinishTime int64
	TotalTime time.Duration
}