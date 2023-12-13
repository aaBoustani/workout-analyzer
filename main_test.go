package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleAnalysis(t *testing.T) {
	requestBody := []byte(`[{"distance": 10000, "time": 3600, "timestamp":"` + getDate(0) + `"}]`)
	req, err := http.NewRequest("POST", "/analyse?nweeks=1", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handleAnalysis(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handleAnalysis returned %v status code; expected %v", status, http.StatusOK)
	}

	expected := `{"medium_distance":10000,"medium_time":3600,` +
		`"max_distance":10000,"max_time":3600,` +
		`"medium_weekly_distance":10000,"medium_weekly_time":3600,` +
		`"max_weekly_distance":10000,"max_weekly_time":3600}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handleAnalysis returned %v as request body; expected %v", rr.Body.String(), expected)
	}
}

func TestHandleAnalysisShouldFail(t *testing.T) {
	// nweeks = 0
	requestBody := []byte(`[{"distance": 10000, "time": 3600, "timestamp":"` + getDate(1) + `"}]`)
	req, err := http.NewRequest("POST", "/analyse?nweeks=0", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handleAnalysis(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handleAnalysis returned %v status code; expected %v", status, http.StatusOK)
	}

	expected := `{"Error": "Number of weeks must be a valid positive number"}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handleAnalysis returned %v as request body; expected %v", rr.Body.String(), expected)
	}

	// nweeks is missing
	req, err = http.NewRequest("POST", "/analyse", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	handleAnalysis(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handleAnalysis returned %v status code; expected %v", status, http.StatusOK)
	}

	expected = `{"Error": "Number of weeks must be a valid positive number"}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handleAnalysis returned %v as request body; expected %v", rr.Body.String(), expected)
	}

	// Broken JSON
	requestBody = []byte(`[{"distance": 10000, "time": 3600, "timestamp":"` + getDate(1) + `"]`)
	req, err = http.NewRequest("POST", "/analyse", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	handleAnalysis(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handleAnalysis returned %v status code; expected %v", status, http.StatusOK)
	}

	expected = `{"Error": "Invalid JSON"}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handleAnalysis returned %v as request body; expected %v", rr.Body.String(), expected)
	}
}
