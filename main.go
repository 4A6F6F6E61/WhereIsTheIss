package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type APIData struct {
	Name       string  `json:"name"`
	Id		   float64 `json:"id"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	Altitude   float64 `json:"altitude"`
	Velocity   float64 `json:"velocity"`
	Visibility float64 `json:"visibility"`
	Footprint  float64 `json:"footprint"`
	Timestamp  float64 `json:"timestamp"`
	Daynum     float64 `json:"daynum"`
	Solar_lat  float64 `json:"solar_lat"`
	Solar_lon  float64 `json:"solar_lon"`
	Units      string  `json:"units"`
}

func main() {
	println("[Log] Searching for the ISS...")
	la, lo := getLocation()
	println("[Log] Location found!")
	time.Sleep(400000000)
	fmt.Printf("\nLatitude:  %f\nLongitude: %f\n", la, lo)
}

func getLocation() (latitude, longitude float64) {
    resp, err := http.Get("https://api.wheretheiss.at/v1/satellites/25544")
    if err != nil {
        println("[Error] Could not get ISS location")
    }
    defer resp.Body.Close()
    bodyBytes, _ := ioutil.ReadAll(resp.Body)
    var data APIData
    json.Unmarshal(bodyBytes, &data)
    //fmt.Printf("API Response as struct %+v\n", data)
	latitude, longitude = data.Latitude, data.Longitude
	return
}