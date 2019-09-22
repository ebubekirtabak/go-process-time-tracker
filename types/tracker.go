package types

import "time"

type Tracker struct {
	Name string 			`json:"name"`
	FirstStartTime int64	`json:"firstStartTime"`
	StartTime int64			`json:"startTime"`
	FinishTime int64		`json:"finishTime"`
	PauseStartTime int64 	`json:"pauseStartTime"`
	TotalTime time.Duration `json:"totalTime"`
	Subtrackers []SubTracker	`json:"subtrackers"`
}

type SubTracker struct {
	SubTrackerName string `json:"delayName"`
	StartTime int64  `json:"startTime"`
	FinishTime int64 `json:"finishTime"`
	TotalTime time.Duration  `json:"totalTime"`
}