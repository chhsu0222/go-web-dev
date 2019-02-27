package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// go to this website to convert JSON to Go struct
// https://mholt.github.io/json-to-go/

// You can choose to only unmarshal some of the json data
// Create a data structure that only has fields for some of the data
type city struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
	City      string  `json:"city"`
}

type cities []city

func main() {
	var data cities
	recvd := `[{"precision":"zip","lat":37.7668,"lon":-122.3959,"Address":"","city":"SAN FRANCISCO","State":"CA","Zip":"94107","Country":"US"},{"precision":"zip","lat":37.371991,"lon":-122.02602,"Address":"","city":"SUNNYVALE","State":"CA","Zip":"94085","Country":"US"}]`
	err := json.Unmarshal([]byte(recvd), &data)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(data)
	fmt.Println(data[1].City)
}
