package measure

import "fmt"
import (
	"go-excute-time-measure/types"
	"time"
)

func Text() {
	fmt.Println("sfsd")
}

var excutes []types.Measure

func StartTime(name string) {
	var measure = types.Measure{}
	measure.Name = name
	measure.StartTime = time.Now().Unix()
	excutes = append(excutes, measure)
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
		out := time.Time{}.Add(diff)
		fmt.Println(time1)
		fmt.Println(time2)
		fmt.Println(out.Format("15:04:05"))
	}
}

func indexOf(element string, data []types.Measure) (int) {
	for k, v := range data {
		fmt.Println(k)
		fmt.Println(v)
		if v.Name == element {
			return k
		}
	}
	return -1    //not found.
}

