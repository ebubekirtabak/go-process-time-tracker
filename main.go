package main

import (
	"go-excute-time-measure/tracker"
	"fmt"
	"time"
	"net/http"
)

func main() {
	tracker.StartTracker("main")
	tracker.StartTracker("sum")
	var sum = 1
	for  i := 0; i < 10000000000; i++ {
		sum = i * sum * sum * i
	}

	fmt.Println(tracker.FinishTracker("sum"))
	tracker.GetExecuteTime("sum")
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
	}

	tracker.StartSubTracker("main", "firstDelay")
	time.Sleep(5 * time.Second)
	tracker.FinishSubTracker("main", "firstDelay")
	time.Sleep(5 * time.Second)
	tracker.StartSubTracker("main", "lastDelay")
	time.Sleep(5 * time.Second)
	tracker.FinishSubTracker("main", "lastDelay")
	time.Sleep(30 * time.Second)
	tracker.StartSubTracker("main", "imageRequest")
	imageResponse, err := http.Get("https://images.pexels.com/photos/2217365/pexels-photo-2217365.jpeg?auto=compress&cs=tinysrgb&dpr=2&w=500")
	if imageResponse.StatusCode == 200 {
		fmt.Printf("StatusCode: %d \n" , imageResponse.StatusCode)
	}
	tracker.FinishSubTracker("main", "imageRequest")
	tracker.GetExecuteTime("main")
	excute, err := tracker.GetExecuteTimeJson("main")
	if err == nil {
		fmt.Println(string(excute))
	}

	excutes, err := tracker.GetAllExecuteTimeJson()
	if err == nil {
		fmt.Println(string(excutes))
	}
}