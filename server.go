package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handleHomeRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home Route")
}

type Response struct {
	Message string
	Items   []int
}

func sendData(w http.ResponseWriter, r *http.Request) {

	// Create a response object
	response := Response{
		Message: "HELLO FROM BACKEND",
		Items:   []int{1, 2, 3},
	}

	// Encode the response struct to JSON and write it to the response writer
	err := json.NewEncoder(w).Encode(response)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error encoding JSON: %v", err), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", handleHomeRoute)
	http.HandleFunc("/data", sendData)

	fmt.Println("Starting Go server on :8080...")
	http.ListenAndServe(":8080", nil)
}
