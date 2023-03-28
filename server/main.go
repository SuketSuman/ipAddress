package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ApiResponse struct {
	IPv string `json:"ipv"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Request-Method", "*")
		w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		// Retrieve IPv6 address from ip6.me API
		resp, err := http.Get("http://ip6.me/api/")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Parse response body and create API response
		// body1 := strings.Split(string(body), "Pv6,")
		// body2 := strings.Split(string(body1[1]), ",v1")
		apiResponse := ApiResponse{IPv: string(body)}

		// Marshal API response to JSON
		jsonResponse, err := json.Marshal(apiResponse)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set response content type to JSON
		w.Header().Set("Content-Type", "application/json")

		// Write JSON response to response body
		w.Write(jsonResponse)
	})

	http.ListenAndServe(":8080", nil)
}
