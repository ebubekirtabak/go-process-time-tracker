package types

import "time"

type Tracker struct {
	Name 			string 		  `json:"name"`
	StartTime 		time.Time	  `json:"startTime"`
	StartTimeUnix 	int64		  `json:"startTimeUnix"`
	FinishTime 		time.Time	  `json:"finishTime"`
	FinishTimeUnix 	int64		  `json:"finishTimeUnix"`
	PauseStartTime  int64 	 	  `json:"pauseStartTime"`
	TotalTime 		time.Duration `json:"totalTime"`
	Subtrackers 	[]SubTracker  `json:"subtrackers"`
}

type SubTracker struct {
	SubTrackerName string 		  `json:"delayName"`
	StartTime 	   time.Time  	  `json:"startTime"`
	StartTimeUnix  int64  		  `json:"startTimeUnix"`
	FinishTime     time.Time 	  `json:"finishTime"`
	FinishTimeUnix int64 		  `json:"finishTimeUnix"`
	TotalTime      time.Duration  `json:"totalTime"`
}