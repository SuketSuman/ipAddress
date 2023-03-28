package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Location struct {
	CountryName string  `json:"country_name"`
	RegionName  string  `json:"region_name"`
	CityName    string  `json:"city_name"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}

func main() {
	http.HandleFunc("/location", getLocation)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getLocation(w http.ResponseWriter, r *http.Request) {
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.RemoteAddr
	}

	url := fmt.Sprintf("https://ipapi.co/%s/json/", ip)

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var loc Location
	err = json.NewDecoder(resp.Body).Decode(&loc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(loc)
}
