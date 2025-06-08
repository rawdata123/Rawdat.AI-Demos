package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Flight struct {
	Flight struct {
		Number string `json:"number"`
	} `json:"flight"`
	Departure struct {
		Airport string `json:"airport"`
	} `json:"departure"`
	Arrival struct {
		Airport string `json:"airport"`
	} `json:"arrival"`
	FlightStatus string `json:"flight_status"`
}

type ApiResponse struct {
	Data []Flight `json:"data"`
}

func main() {
	http.HandleFunc("/", flightsHandler)
	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func flightsHandler(w http.ResponseWriter, r *http.Request) {
	apiKey := os.Getenv("AVIATIONSTACK_API_KEY")
	if apiKey == "" {
		http.Error(w, "Missing API key", http.StatusInternalServerError)
		return
	}
	url := fmt.Sprintf("http://api.aviationstack.com/v1/flights?access_key=%s&flight_status=active&limit=10", apiKey)

	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		http.Error(w, "Failed to fetch flights", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var apiResp ApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		http.Error(w, "Failed to decode response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Real-Time Flights</h1><ul>")
	for _, flight := range apiResp.Data {
		fmt.Fprintf(w,
			"<li>Flight %s: %s → %s (%s)</li>",
			flight.Flight.Number,
			flight.Departure.Airport,
			flight.Arrival.Airport,
			flight.FlightStatus,
		)
	}
	fmt.Fprint(w, "</ul>")
}

