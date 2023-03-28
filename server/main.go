package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
)

type IPv6Response struct {
	IPAddress string `json:"ip_address"`
}

func main() {
	http.HandleFunc("/ipv6", ipv6Handler)
	http.ListenAndServe(":8080", nil)
}

func ipv6Handler(w http.ResponseWriter, r *http.Request) {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Error: ", err.Error())
		return
	}

	ipv6 := net.ParseIP(ip)
	if ipv6 == nil || ipv6.To4() != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Error: IP address is not IPv6")
		return
	}

	response := IPv6Response{IPAddress: ipv6.String()}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error: failed to encode response as JSON")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
