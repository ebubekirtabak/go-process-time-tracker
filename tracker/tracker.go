package tracker

import "fmt"
import (
	"go-excute-time-measure/types"
	"time"
	"encoding/json"
)

var excutes []types.Tracker

func StartTracker(name string) {
	var measure = types.Tracker{}
	measure.Name = name
	measure.FirstStartTime = time.Now().Unix()
	measure.StartTime = measure.FirstStartTime
	excutes = append(excutes, measure)
}

func FinishSubTracker(name string, subTrackerName string) {
	index := indexOf(name, excutes)
	if index > -1 {
		var measure= excutes[index]
		measure.StartTime = time.Now().Unix()
		subTrackerIndex := indexOfSubTracker(subTrackerName, measure)
		if subTrackerIndex > -1 {
			subTracker := measure.Subtrackers[subTrackerIndex]
			subTracker.StartTime = measure.FinishTime
			subTracker.FinishTime = measure.StartTime
			startTime := time.Unix(subTracker.StartTime, 0)
			finishTime := time.Unix(subTracker.FinishTime, 0)
			diff := finishTime.Sub(startTime)
			subTracker.TotalTime += diff
			measure.Subtrackers[subTrackerIndex] = subTracker
		}

		measure.FinishTime = 0
		excutes[index] = measure
	} else {
		StartTracker(name)
	}
}

func StartSubTracker(name string, subTrackerName string) {
	index := indexOf(name, excutes)
	if index > -1 {
		var measure= excutes[index]
		measure.Name = name
		measure.FinishTime = time.Now().Unix()
		subTracker := types.SubTracker{}
		subTracker.StartTime = measure.FinishTime
		subTracker.SubTrackerName = subTrackerName
		measure.Subtrackers = append(measure.Subtrackers, subTracker)
		time1 := time.Unix(measure.StartTime, 0)
		time2 := time.Unix(measure.FinishTime, 0)
		diff := time2.Sub(time1)
		measure.TotalTime += diff
		out := time.Time{}.Add(diff)
		fmt.Println(out.Format("15:04:05"))
		excutes[index] = measure
	}
}

func FinishTracker(name string) types.Tracker {
	index := indexOf(name, excutes)
	if index > -1 {
		var measure = excutes[index]
		measure.Name = name
		measure.FinishTime = time.Now().Unix()
		excutes[index] = measure
		return measure
	}

	return types.Tracker{}
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

	return nil, fmt.Errorf("Execute not found: %s", name)
}

func GetAllExecuteTimeJson() ([]byte, error)  {

	excutesJson, err := json.Marshal(excutes)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return excutesJson, nil
}


func indexOf(element string, data []types.Tracker) (int) {
	for k, v := range data {
		if v.Name == element {
			return k
		}
	}
	return -1 //not found.
}

func indexOfSubTracker(element string, measure types.Tracker) (int) {
	for k, v := range measure.Subtrackers {
		if v.SubTrackerName == element {
			return k
		}
	}
	return -1 //not found.
}


