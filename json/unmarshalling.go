package main

import (
	"encoding/json"
	"fmt"
)

type SensorReading struct {
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	Time     string `json:"time"`
}

func main() {
	jsonString := `{ "name": "battery sensor", "capacity": 40, "time": "2019-01-21T19:07:28Z" }`

	var reading SensorReading
	var reading2 map[string]interface{}
	_ = json.Unmarshal([]byte(jsonString), &reading)
	_ = json.Unmarshal([]byte(jsonString), &reading2)
	fmt.Printf("%+v\n", reading)
	fmt.Printf("%v\n", reading)
	fmt.Printf("%+v\n", reading2)
	fmt.Printf("%v\n", reading2["name"])
}
