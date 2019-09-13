package measure

import "fmt"
import (
	"go-excute-time-measure/types"
	"time"
	"encoding/json"
)

var excutes []types.Measure

func StartTime(name string) {
	var measure = types.Measure{}
	measure.Name = name
	measure.FirstStartTime = time.Now().Unix()
	measure.StartTime = measure.FirstStartTime
	excutes = append(excutes, measure)
}

func ResumeTime(name string) {
	index := indexOf(name, excutes)
	if index > -1 {
		var measure= excutes[index]
		measure.StartTime = time.Now().Unix()
		delayLog := types.DelayLog{}
		delayLog.StartTime = measure.FinishTime
		delayLog.FinishTime = measure.StartTime

		measure.Delaylogs = append(measure.Delaylogs, delayLog)
		measure.FinishTime = 0
		excutes[index] = measure
	} else {
		StartTime(name)
	}
}

func PauseTime(name string) {
	index := indexOf(name, excutes)
	if index > -1 {
		var measure= excutes[index]
		measure.Name = name
		measure.FinishTime = time.Now().Unix()
		time1 := time.Unix(measure.StartTime, 0)
		time2 := time.Unix(measure.FinishTime, 0)
		diff := time2.Sub(time1)
		measure.TotalTime += diff
		out := time.Time{}.Add(diff)
		fmt.Println(out.Format("15:04:05"))
		excutes[index] = measure
	}
}

func FinishTime(name string) types.Measure {
	index := indexOf(name, excutes)
	if index > -1 {
		var measure = excutes[index]
		measure.Name = name
		measure.FinishTime = time.Now().Unix()
		excutes[index] = measure
		return measure
	}

	return types.Measure{}
}

func GetExecuteTime(name string) {
	index := indexOf(name, excutes)
	if index > -1 {
		var measure = excutes[index]
		time1 := time.Unix(measure.StartTime, 0)
		time2 := time.Unix(measure.FinishTime, 0)
		diff := time2.Sub(time1)
		measure.TotalTime += diff
		out := time.Time{}.Add(measure.TotalTime)
		fmt.Println(out.Format("15:04:05"))
	}
}

func GetExecuteTimeJson(name string) ([]byte, error)  {
	index := indexOf(name, excutes)
	if index > -1 {
		var measure = excutes[index]
		measureJson, err := json.Marshal(measure)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		return measureJson, nil
	}
}

func indexOf(element string, data []types.Measure) (int) {
	for k, v := range data {
		if v.Name == element {
			return k
		}
	}
	return -1 //not found.
}

