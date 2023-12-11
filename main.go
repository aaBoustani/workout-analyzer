package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func handleAnalysis(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var workouts []Workout
	err := json.NewDecoder(r.Body).Decode(&workouts)
	if err != nil {
		http.Error(w, `{"Error": "Invalid JSON"}`, http.StatusBadRequest)
		return
	}

	nWeeksStr := r.URL.Query().Get("nweeks")
	nWeeks, err := strconv.Atoi(nWeeksStr)
	if err != nil || nWeeks == 0 {
		http.Error(w, `{"Error": "Number of weeks must be a valid positive number"}`, http.StatusBadRequest)
		return
	}

	response := Analyze(workouts, nWeeks)
	json.NewEncoder(w).Encode(response)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/analyse", handleAnalysis).Methods("POST")

	log.Println("Starting the server on port 3000")
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal(err)
	}
}
