
# GOTIMELOG

Go Time Log is time measure library for your function time.

**WHY ?**
-
We can want measure execution time for our functions.

**Features**
- 
* Create multiple tracker.
* Pause Tracker.
* Resume Tracker.

**Usage**
-

Start a single tracker:
```GO
measure.StartTime("$TRACKER_NAME")
```

You can **Pause** and **Resume** tracker.
Pause tracker:

```GO
measure.PauseTime("$TRACKER_NAME", "$DELAY_POINT_NAME")
```

**Resume**

````GO
measure.ResumeTime("$DELAY_POINT_NAME", "$DELAY_POINT_NAME")
````

Test Code:
```GO
measure.StartTime("$TRACKER_NAME")
var sum = 1
for  i := 0; i < 100000000000; i++ {
	sum = i * sum * sum * i
}
measure.FinishTime("$TRACKER_NAME")
mainTime := measure.GetExecuteTime("$TRACKER_NAME")
// OR
mainTimeJson, err := measure.GetExecuteTimeJson("$TRACKER_NAME")
```

Response: 

```json
{
  "name": "$TRACKER_NAME",
  "firstStartTime": 1568569495,
  "startTime": 1568569520,
  "finishTime": 1568569550,
  "pauseStartTime": 0,
  "totalTime": 15000000000,
  "delayLogs": [
    {
      "delayName": "firstDelay",
      "startTime": 1568569505,
      "finishTime": 1568569510,
      "totalTime": 5000000000
    },
    {
      "delayName": "lastDelay",
      "startTime": 1568569515,
      "finishTime": 1568569520,
      "totalTime": 5000000000
    }
  ]
}
```

**Contributing**
-
Contributions are welcome! Fork this repo and add your changes and submit a PR.

If you would like to fix a bug, add a feature or provide feedback you can do so in the issues section.
