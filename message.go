package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Message is used to form JSON data
type Message struct {
	Year   int
	Month  string
	Date   int
	Hour   int
	Minute int
	Second int
	Days   int
}

func getTime() float64 {
	loc, _ := time.LoadLocation("PRC")
	loveDate := time.Date(2017, 8, 26, 1, 0, 0, 0, loc)
	diff := time.Now().Sub(loveDate)
	return diff.Hours() / 24
}

func buildMsg() Message {
	msg := Message{time.Now().Year(), time.Now().Month().String(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), int(getTime())}
	return msg
}

func serveHTTP(w http.ResponseWriter, r *http.Request) {
	msg := buildMsg()
	log.Printf("msg: %v", msg)
	jsonMsg, _ := json.Marshal(msg)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonMsg)
}
