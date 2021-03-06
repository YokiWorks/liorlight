package main

import (
	"fmt"
	"time"

	"github.com/yryz/ds18b20"
)

func init() {
	go startWater()
}

func startWater() {

	sensors, err := ds18b20.Sensors()
	if err != nil {
		fmt.Println(err)
	}
	if sensors == nil {
		fmt.Println("water sensor not found")
		return
	}

	for {

		for _, sensor := range sensors {
			t, err := ds18b20.Temperature(sensor)
			if err == nil {
				waterTemp.Set(t)
				fmt.Printf("sensor: %s temperature: %.2f°C\n", sensor, t)
			} else {
				logger.Error(err)
			}
		}
		time.Sleep(60 * time.Second)
	}
}
