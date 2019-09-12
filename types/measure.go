package types

import "time"

type Measure struct {
	Name string 			`json:"name"`
	FirstStartTime int64	`json:"firstStartTime"`
	StartTime int64			`json:"startTime"`
	FinishTime int64		`json:"finishTime"`
	PauseStartTime int64 	`json:"pauseStartTime"`
	TotalTime time.Duration `json:"totalTime"`
	Delaylogs []DelayLog	`json:"delayLogs"`
}

type DelayLog struct {
	StartTime int64  `json:"startTime"`
	FinishTime int64 `json:"finishTime"`
	TotalTime int64  `json:"totalTime"`
}