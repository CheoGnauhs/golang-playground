package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/msg", serveHTTP)
	http.HandleFunc("/aqi", getAqi)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
