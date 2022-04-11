package main

import (
	"1001_exercise/geo"
	"fmt"
	"log"
)

func main() {
	coordinates := geo.Coordinates{}
	err := coordinates.SetLatitude(36.42)
	if err != nil {
		log.Fatal(err)
	}
	err = coordinates.SetLongitude(-122.00)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(coordinates.Latitude())
	fmt.Println(coordinates.Longitude())

	landmark := geo.Landmark{}
	err = landmark.SetName("AnyWhere")
	if err != nil {
		log.Fatal(err)
	}
	err = landmark.SetLatitude(37.22)
	if err != nil {
		log.Fatal(err)
	}
	err = landmark.SetLongitude(120.12)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(landmark.Name())
	fmt.Println(landmark.Latitude())
	fmt.Println(landmark.Longitude())
}
