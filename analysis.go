package main

import (
	"log"
	"reflect"
	"time"
)

func calculateTotal(workouts []Workout, field string) int {
	total := 0
	for _, w := range workouts {
		total += int(reflect.ValueOf(w).FieldByName(field).Int())
	}
	return total
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxByField(workouts []Workout, field string) int {
	var max int = 0
	for _, w := range workouts {
		value := int(reflect.ValueOf(w).FieldByName(field).Int())
		max = maxInt(max, value)
	}
	return max
}

func weeklyMaxByField(weeklyData map[int][]Workout, field string) int {
	max := 0
	for _, weekData := range weeklyData {
		weekly := calculateTotal(weekData, field)
		max = maxInt(max, weekly)
	}
	return max
}

func passesThreshold(thresholdYear, thresholdWeek, year, week int) bool {
	return thresholdYear < year || (thresholdYear == year && thresholdWeek < week)
}

func filterAndGroupByWeek(workouts []Workout, nWeeks int) ([]Workout, map[int][]Workout) {
	// Threshold to filter data that are older than n weeks.
	thresholdYear, thresholdWeek := time.Now().AddDate(0, 0, -nWeeks*7).ISOWeek()
	layout := "2006-01-02T15:04:05.000000Z"

	workoutsMap := make(map[int][]Workout)
	var filteredWorkouts []Workout

	for _, w := range workouts {
		t, err := time.Parse(layout, w.Timestamp)
		if err != nil {
			// Note: This could also be log.Fatal() depending on how tolerant we want to be with the failures.
			log.Println("Error parsing timestamp: ", err)
			continue
		}
		year, week := t.ISOWeek()
		if passesThreshold(thresholdYear, thresholdWeek, year, week) {
			workoutsMap[year*100+week] = append(workoutsMap[year*100+week], w)
			filteredWorkouts = append(filteredWorkouts, w)
		}
	}
	return filteredWorkouts, workoutsMap
}

func Analyze(workouts []Workout, nWeeks int) Analysis {
	workouts, weeklyData := filterAndGroupByWeek(workouts, nWeeks)

	// Default value. This means that there aren't any workouts for the past n weeks.
	if len(workouts) == 0 {
		return Analysis{
			MedDistance:       0,
			MedTime:           0,
			MaxDistance:       0,
			MaxTime:           0,
			MedWeeklyDistance: 0,
			MedWeeklyTime:     0,
			MaxWeeklyDistance: 0,
			MaxWeeklyTime:     0,
		}
	}

	totalDistance := calculateTotal(workouts, "Distance")
	totalTime := calculateTotal(workouts, "Time")

	return Analysis{
		MedDistance:       totalDistance / len(workouts),
		MedTime:           totalTime / len(workouts),
		MaxDistance:       maxByField(workouts, "Distance"),
		MaxTime:           maxByField(workouts, "Time"),
		MedWeeklyDistance: totalDistance / nWeeks,
		MedWeeklyTime:     totalTime / nWeeks,
		MaxWeeklyDistance: weeklyMaxByField(weeklyData, "Distance"),
		MaxWeeklyTime:     weeklyMaxByField(weeklyData, "Time"),
	}
}
