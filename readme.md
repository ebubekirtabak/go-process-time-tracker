
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
    "startTime": "2019-10-08T22:28:14.740610487+03:00",
    "startTimeUnix": 1570562956,
    "finishTime": "2019-10-08T22:29:15.869431745+03:00",
    "finishTimeUnix": 0,
    "pauseStartTime": 0,
    "totalTime": 51000000000,
    "subtrackers": [
      {
        "delayName": "firstDelay",
        "startTime": "2019-10-08T22:28:30.853985579+03:00",
        "startTimeUnix": 1570562910,
        "finishTime": "2019-10-08T22:28:35.859056535+03:00",
        "finishTimeUnix": 1570562915,
        "totalTime": 5000000000
      },
      {
        "delayName": "lastDelay",
        "startTime": "2019-10-08T22:28:40.861031205+03:00",
        "startTimeUnix": 1570562920,
        "finishTime": "2019-10-08T22:28:45.86572914+03:00",
        "finishTimeUnix": 1570562925,
        "totalTime": 5000000000
      },
      {
        "delayName": "imageRequest",
        "startTime": "2019-10-08T22:29:15.869431745+03:00",
        "startTimeUnix": 1570562955,
        "finishTime": "2019-10-08T22:29:16.214073394+03:00",
        "finishTimeUnix": 1570562956,
        "totalTime": 1000000000
      }
    ]
  },
  {
    "name": "sum",
    "startTime": "2019-10-08T22:28:14.740613261+03:00",
    "startTimeUnix": 1570562894,
    "finishTime": "2019-10-08T22:28:20.816077878+03:00",
    "finishTimeUnix": 1570562900,
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
