package main

import (
	"fmt"
	"net/http"
)

func handleHomeRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home Route")
}

func main() {
	http.HandleFunc("/", handleHomeRoute)
	fmt.Println("Starting Go server on :8080...")
	http.ListenAndServe(":8080", nil)
}
