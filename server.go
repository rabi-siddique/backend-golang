package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handleHomeRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home Route")
}

type Response struct {
	Message string
	Items   []int
}

type RequestData struct {
	Name     string
	LastName string
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

func submitData(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var data RequestData
		decoder := json.NewDecoder(r.Body)

		err := decoder.Decode(&data)
		if err != nil {
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
			return
		}

		if data.Name == "" || data.LastName == "" {
			log.Println("Name or LastName is missing")
			http.Error(w, "Name or LastName is misisng", http.StatusBadRequest)
			return
		}

		fmt.Println("Data Received:", data)
		fmt.Println("Name", data.Name)
		fmt.Println("LastName", data.LastName)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Data successfully received"))
	} else {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}
}

func main() {
	http.HandleFunc("/", handleHomeRoute)
	http.HandleFunc("/data", sendData)
	http.HandleFunc("/submit-data", submitData)

	fmt.Println("Starting Go server on :8080...")
	http.ListenAndServe(":8080", nil)
}
