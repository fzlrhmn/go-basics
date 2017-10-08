package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
)

type Coordinate struct {
	LatitudeE7  int
	LongitudeE7 int
	TimestampMs string
}

type Locations struct {
	Locations []Coordinate
}

type Coord struct {
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Time      time.Time `json:"time,omitempty"`
}

type MyLoc struct {
	Coords []Coord `json:"coordinates"`
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func (loc *MyLoc) addCoord(coord Coord) {
	loc.Coords = append(loc.Coords, coord)
}

func main() {
	j, err := ioutil.ReadFile("storage/location.json")
	checkErr(err)

	coordinates := Locations{}
	json.Unmarshal(j, &coordinates)

	locs := MyLoc{}
	for _, v := range coordinates.Locations {
		// coordinates
		lat := float64(v.LatitudeE7) / 10000000
		lon := float64(v.LongitudeE7) / 10000000

		// timestamp
		i, err := strconv.ParseInt(v.TimestampMs, 10, 64)
		checkErr(err)
		tm := time.Unix(0, i*int64(time.Millisecond))

		locs.addCoord(Coord{Latitude: lat, Longitude: lon, Time: tm})
	}

	jsonResp, _ := json.Marshal(locs)

	fmt.Println(string(jsonResp))
}
