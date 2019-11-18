package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

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
	w.Write(resBody)
}
