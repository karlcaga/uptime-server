package main

import (
	"log"
	"net/http"
)

func main() {
	log.SetPrefix("Uptime Server: ")

	http.HandleFunc("POST /{nodeName}", handleMetric)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleMetric(w http.ResponseWriter, r *http.Request) {
	nodeName := r.PathValue("nodeName")
	log.Println(nodeName)
}
