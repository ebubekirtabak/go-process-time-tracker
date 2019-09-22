package types

import "time"

type Tracker struct {
	Name string 			`json:"name"`
	StartTime time.Time		`json:"startTime"`
	StartTimeUnix int64		`json:"startTimeUnix"`
	FinishTime int64		`json:"finishTime"`
	PauseStartTime int64 	`json:"pauseStartTime"`
	TotalTime time.Duration `json:"totalTime"`
	Subtrackers []SubTracker	`json:"subtrackers"`
}

type SubTracker struct {
	SubTrackerName string 		  `json:"delayName"`
	StartTime 	   time.Time  	  `json:"startTime"`
	StartTimeUnix  int64  		  `json:"startTime"`
	FinishTime     int64 		  `json:"finishTime"`
	TotalTime      time.Duration  `json:"totalTime"`
}