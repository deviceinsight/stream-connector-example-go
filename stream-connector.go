package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Datapoint struct {
	Key       string
	Value     string
	Timestamp string
}

type Event struct {
	EventKey  string
	Priority  int
	Timestamp string
}

type GeoPosition struct {
	Latitude  float32
	Longitude float32
	Timestamp string
}

func writeJSON(d interface{}) {
	b, err := json.Marshal(d)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
}

func main() {
	interval := 5
	for {
		datapoint := Datapoint{
			Key:       "test-key",
			Value:     strconv.Itoa(rand.Intn(100)),
			Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),
		}
		writeJSON(datapoint)
		fmt.Println()
		time.Sleep(time.Duration(interval) * time.Second)
	}
}
