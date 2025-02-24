package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "pong")
}

func challenge1Handler(w http.ResponseWriter, r *http.Request) {
	var array []int
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&array); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Write the code for Challenge 1 here
	result := 0

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func challenge2Handler(w http.ResponseWriter, req *http.Request) {
	var array []int
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&array); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Write the code for Challenge 2 here
	result := 0
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func challenge3Handler(w http.ResponseWriter, req *http.Request) {
	var array []int
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&array); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Write the code for Challenge 3 here
	result := 0

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func main() {
	http.HandleFunc("/ping", handler)
	http.HandleFunc("/chalange-1", challenge1Handler)
	http.HandleFunc("/chalange-2", challenge2Handler)
	http.HandleFunc("/chalange-3", challenge3Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
