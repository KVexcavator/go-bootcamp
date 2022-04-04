package main

import (
	"830_export/geo/geo"
	"fmt"
)

func main() {
	location := geo.Coordinates{Latitude: 37.42, Longitude: -127.03}
	fmt.Println("Latitude:", location.Latitude)
	fmt.Println("Longitude:", location.Longitude)
}
