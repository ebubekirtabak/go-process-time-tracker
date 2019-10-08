
# GOTIMETRACKER

Go Time Tracker is time measure library for your function time.

**WHY ?**
-
We can want measure execution time for our functions.

**Features**
- 
* Create multiple tracker.
* Sub Tracker.

**Usage**
-

Start a single tracker:
```GO
tracker.StartTime("$TRACKER_NAME")
```

You can create multiple **Sub Tracker**.

Start Sub Tracker:

```GO
tracker.StartSubTracker("$TRACKER_NAME", "$SUB_TRACKER_NAME")
```

**Resume**

````GO
tracker.FinishSubTracker("$TRACKER_NAME", "$SUB_TRACKER_NAME")
````

**Finish**

```GO
tracker.FinishTime("$TRACKER_NAME")
```

Test Code:
```GO
tracker.StartTime("$TRACKER_NAME")
var sum = 1
for  i := 0; i < 100000000000; i++ {
	sum = i * sum * sum * i
}
tracker.FinishTime("$TRACKER_NAME")
mainTime := measure.GetExecuteTime("$TRACKER_NAME")
// OR
mainTimeJson, err := tracker.GetExecuteTimeJson("$TRACKER_NAME")
```

Response: 

```json
[
  {
    "name": "main",
    "startTime": "2019-09-25T21:33:33.092833996+03:00",
    "startTimeUnix": 1569436471,
    "finishTime": "2019-09-25T21:34:31.302779044+03:00",
    "finishTimeUnix": 0,
    "pauseStartTime": 0,
    "totalTime": 48000000000,
    "subtrackers": [
      {
        "delayName": "firstDelay",
        "startTime": "2019-09-25T21:33:46.290234007+03:00",
        "startTimeUnix": 1569436426,
        "finishTime": "0001-01-01T00:00:00Z",
        "finishTimeUnix": 1569436431,
        "totalTime": 5000000000
      },
      {
        "delayName": "lastDelay",
        "startTime": "2019-09-25T21:33:56.296966325+03:00",
        "startTimeUnix": 1569436436,
        "finishTime": "0001-01-01T00:00:00Z",
        "finishTimeUnix": 1569436441,
        "totalTime": 5000000000
      },
      {
        "delayName": "imageRequest",
        "startTime": "2019-09-25T21:34:31.302779044+03:00",
        "startTimeUnix": 1569436471,
        "finishTime": "0001-01-01T00:00:00Z",
        "finishTimeUnix": 1569436471,
        "totalTime": 0
      }
    ]
  },
  {
    "name": "sum",
    "startTime": "2019-09-25T21:33:33.092836126+03:00",
    "startTimeUnix": 1569436413,
    "finishTime": "2019-09-25T21:33:36.261081316+03:00",
    "finishTimeUnix": 1569436416,
    "pauseStartTime": 0,
    "totalTime": 0,
    "subtrackers": null
  }
]
```

**Contributing**
-
Contributions are welcome! Fork this repo and add your changes and submit a PR.

If you would like to fix a bug, add a feature or provide feedback you can do so in the issues section.
