package main

import (
	"go-excute-time-measure/measure"
	"fmt"
	"time"
)

func main() {
	measure.StartTime("main")
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
	}
	measure.PauseTime("main")
	time.Sleep(5 * time.Second)
	measure.ResumeTime("main")
	time.Sleep(5 * time.Second)
	fmt.Println(measure.FinishTime("main"))
	measure.GetExecuteTime("main")
}