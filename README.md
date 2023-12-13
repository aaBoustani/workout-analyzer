# Workout Analyzer
A simple server in Golang with a single endpoint: `POST /analyse?nweeks=`

## Clone the project
```
$ git clone https://go.googlesource.com/example
$ cd workout-analyzer
```

## Usage
```
$ make run
```
This command will install dependencies, build the project, then run the binary object, which will open a server on `localhost:3000`.

## Example
### Endpoint
`POST /analyse?nweeks=6`  
Set the number of weeks over which you want to see the statistics.
### Body
```
[{
    "distance": 10000,
    "time": 3600,
    "timestamp":"2023-12-07T13:43:28.073909Z"
}]
```
Update `timestamp` to be within the same week as the day you run the request. The week starts on Monday.
### Response
```
{
    "medium_distance": 10000,
    "medium_time": 3600,
    "max_distance": 10000,
    "max_time": 3600,
    "medium_weekly_distance": 1666,
    "medium_weekly_time": 600,
    "max_weekly_distance": 10000,
    "max_weekly_time": 3600
}
```
## Dependencies
This service uses an external library, which is saved in `go.mod`. To install the dependencies (outside of usage), run the following command:
```
$ make install
```
Which installs all dependencies in `go.mod`.

## Test
The project contains testing to test the analysis of the workout data as well as the controller `handleAnalysis` to make sure that `nweeks` is a valid positive number and the input is in fact a JSON object.

```
$ make test
```

Expected output:
```
=== RUN   TestCalculateTotal
--- PASS: TestCalculateTotal (0.00s)
=== RUN   TestMaxInt
--- PASS: TestMaxInt (0.00s)
=== RUN   TestMaxByField
--- PASS: TestMaxByField (0.00s)
=== RUN   TestWeeklyMaxByField
--- PASS: TestWeeklyMaxByField (0.00s)
=== RUN   TestPassesThreshold
--- PASS: TestPassesThreshold (0.00s)
=== RUN   TestFilterAndGroupByWeek
--- PASS: TestFilterAndGroupByWeek (0.00s)
=== RUN   TestAnalyze
--- PASS: TestAnalyze (0.00s)
=== RUN   TestHandleAnalysis
--- PASS: TestHandleAnalysis (0.00s)
=== RUN   TestHandleAnalysisShouldFail
--- PASS: TestHandleAnalysisShouldFail (0.00s)
PASS
ok      command-line-arguments  0.614s
```