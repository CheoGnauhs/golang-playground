package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Weather is used to store json data
type Weather struct {
	Status string
	Data   WeatherData
}

// WeatherData is used to parse aqi data
type WeatherData struct {
	Aqi         int
	Dominentpol string
	Time        Time
}

// Time is used to parse time
type Time struct {
	LocalTime string `json:"s"`
}

func getToken() map[string]string {
	fileValue, err := ioutil.ReadFile("aqi-token.json")
	if err != nil {
		log.Fatal(err)
	}
	var jsonMap map[string]string
	err = json.Unmarshal(fileValue, &jsonMap)
	if err != nil {
		log.Fatal(err)
	}
	return jsonMap
}

func getAqi(w http.ResponseWriter, r *http.Request) {
	jsonStr := getToken()
	url := fmt.Sprintf("http://api.waqi.info/feed/shanghai/?token=%s", jsonStr["token"])
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Request 200 OK")
	var weather Weather
	err = json.Unmarshal(resBody, &weather)
	if err != nil {
		log.Fatal(err)
	}
	if weather.Status == "ok" {
		jsonString := fmt.Sprintf(`{"time":%s,"dominentpol":%s,"aqi":%d}`, weather.Data.Time.LocalTime, weather.Data.Dominentpol, weather.Data.Aqi)
		fmt.Println(jsonString)
		w.Write([]byte(jsonString))
	}
}
