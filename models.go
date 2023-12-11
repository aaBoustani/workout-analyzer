package main

type Analysis struct {
	MedDistance       int `json:"medium_distance"`
	MedTime           int `json:"medium_time"`
	MaxDistance       int `json:"max_distance"`
	MaxTime           int `json:"max_time"`
	MedWeeklyDistance int `json:"medium_weekly_distance"`
	MedWeeklyTime     int `json:"medium_weekly_time"`
	MaxWeeklyDistance int `json:"max_weekly_distance"`
	MaxWeeklyTime     int `json:"max_weekly_time"`
}

type Workout struct {
	Distance  int    `json:"distance"`
	Time      int    `json:"time"`
	Timestamp string `json:"timestamp"`
}
