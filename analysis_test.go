package main

import (
	"testing"
	"time"
)

type metricsTest struct {
	workouts []Workout
	expected Workout
}

var defaultAnalysis = Analysis{
	MedDistance:       0,
	MedTime:           0,
	MaxDistance:       0,
	MaxTime:           0,
	MedWeeklyDistance: 0,
	MedWeeklyTime:     0,
	MaxWeeklyDistance: 0,
	MaxWeeklyTime:     0,
}

func getDate(nWeeks int) string {
	return time.Now().AddDate(0, 0, -nWeeks*7).Format("2006-01-02T15:04:05.000000Z")
}

func TestCalculateTotal(t *testing.T) {
	metrics := metricsTest{
		workouts: []Workout{
			{Distance: 5000, Time: 1000},
			{Distance: 7000, Time: 1500},
		},
		expected: Workout{
			Distance: 12000,
			Time:     2500,
		},
	}

	totalDistance := calculateTotal(metrics.workouts, "Distance")
	totalTime := calculateTotal(metrics.workouts, "Time")

	if totalDistance != metrics.expected.Distance {
		t.Errorf("calculateTotal(workouts, \"Distance\") returned %d; expected %d", totalDistance, metrics.expected.Distance)
	}
	if totalTime != metrics.expected.Time {
		t.Errorf("calculateTotal(workouts, \"Time\") returned %d; expected %d", totalTime, metrics.expected.Time)
	}
}

func TestMaxInt(t *testing.T) {
	if maxInt(3, 5) != 5 {
		t.Errorf("maxInt(3, 5) returned %d; expected %d", maxInt(3, 5), 5)
	}
	if maxInt(7, 3) != 7 {
		t.Errorf("maxInt(7, 3) returned %d; expected %d", maxInt(7, 3), 7)
	}
}

func TestMaxByField(t *testing.T) {
	metrics := metricsTest{
		workouts: []Workout{
			{Distance: 5000, Time: 1000},
			{Distance: 7000, Time: 1500},
		},
		expected: Workout{
			Distance: 7000,
			Time:     1500,
		},
	}

	maxDistance := maxByField(metrics.workouts, "Distance")
	maxTime := maxByField(metrics.workouts, "Time")

	if maxDistance != metrics.expected.Distance {
		t.Errorf("maxByField(workouts, \"Distance\") returned %d; expected %d", maxDistance, metrics.expected.Distance)
	}
	if maxTime != metrics.expected.Time {
		t.Errorf("maxByField(workouts, \"Time\") returned %d; expected %d", maxTime, metrics.expected.Time)
	}
}

func TestWeeklyMaxByField(t *testing.T) {
	weeklyData := map[int][]Workout{
		202341: {{Distance: 5000, Time: 1000}, {Distance: 7000, Time: 1500}, {Distance: 11000, Time: 2200}},
		202342: {{Distance: 7000, Time: 1500}, {Distance: 11000, Time: 2300}, {Distance: 3000, Time: 1900}},
	}

	expected := Workout{
		Distance: 23000,
		Time:     5700,
	}

	maxWeeklyDistance := weeklyMaxByField(weeklyData, "Distance")
	maxWeeklyTime := weeklyMaxByField(weeklyData, "Time")

	if maxWeeklyDistance != expected.Distance {
		t.Errorf("weeklyMaxByField(weeklyData, \"Distance\") returned %d; expected %d", maxWeeklyDistance, expected.Time)
	}
	if maxWeeklyTime != expected.Time {
		t.Errorf("weeklyMaxByField(weeklyData, \"Time\") returned %d; expected %d", maxWeeklyTime, expected.Distance)
	}
}

func TestFilterAndGroupByWeek(t *testing.T) {
	workouts := []Workout{
		{Distance: 5000, Time: 1000, Timestamp: getDate(8)},
		{Distance: 7000, Time: 1500, Timestamp: getDate(1)},
		{Distance: 5000, Time: 1000, Timestamp: getDate(5)},
		{Distance: 7000, Time: 1500, Timestamp: getDate(3)},
		{Distance: 5000, Time: 1000, Timestamp: getDate(8)},
		{Distance: 7000, Time: 1500, Timestamp: getDate(2)},
		{Distance: 5000, Time: 1000, Timestamp: getDate(9)},
		{Distance: 7000, Time: 1500, Timestamp: getDate(9)},
	}

	nWeeks := 8
	expected_len := 6
	numWeeks := 5
	getGroupKey := func(nw int) int {
		year, week := time.Now().AddDate(0, 0, -nw*7).ISOWeek()
		return year*100 + week
	}

	filteredWorkouts, weeklyData := filterAndGroupByWeek(workouts, nWeeks)

	if len(filteredWorkouts) != expected_len {
		t.Errorf("filterAndGroupByWeek returned %d workouts; expected %d", len(filteredWorkouts), expected_len)
	}
	if len(weeklyData) != numWeeks {
		t.Errorf("filterAndGroupByWeek returned %d weeks of data; expected %d", len(weeklyData), numWeeks)
	}
	groupKey := getGroupKey(8)
	if len(weeklyData[groupKey]) != 2 {
		t.Errorf("filterAndGroupByWeek returned %d data elements for week %d; expected %d", len(weeklyData[groupKey]), 8, 2)
	}
	for _, value := range []int{1, 3, 5} {
		groupKey = getGroupKey(value)
		if len(weeklyData[groupKey]) != 1 {
			t.Errorf("filterAndGroupByWeek returned %d data elements for week %d; expected %d", len(weeklyData[groupKey]), value, 1)
		}
	}
}

func TestAnalyze(t *testing.T) {
	workouts := []Workout{
		{Distance: 5000, Time: 1000, Timestamp: getDate(8)},
		{Distance: 7000, Time: 1500, Timestamp: getDate(9)},
		{Distance: 3000, Time: 1200, Timestamp: getDate(2)},
		{Distance: 2000, Time: 1300, Timestamp: getDate(2)},
		{Distance: 1000, Time: 900, Timestamp: getDate(4)},
		{Distance: 5000, Time: 2000, Timestamp: getDate(3)},
		{Distance: 7000, Time: 1500, Timestamp: getDate(3)},
		{Distance: 5000, Time: 3000, Timestamp: getDate(2)},
	}

	nWeeks := 1
	stats := Analyze(workouts, nWeeks)
	if stats != defaultAnalysis {
		t.Errorf("analyze returned %+v; expected %+v", stats, defaultAnalysis)
	}

	nWeeks = 4
	stats = Analyze(workouts, nWeeks)
	expected := Analysis{
		MedDistance:       3833,
		MedTime:           1650,
		MaxDistance:       7000,
		MaxTime:           3000,
		MedWeeklyDistance: 5750,
		MedWeeklyTime:     2475,
		MaxWeeklyDistance: 12000,
		MaxWeeklyTime:     5500,
	}
	if stats != expected {
		t.Errorf("analyze returned %+v; expected %+v", stats, expected)
	}
}
