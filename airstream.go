package main

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func init() {
	ticker := time.NewTicker(120 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				go pool()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func pool() {
	// Read DHT11 sensor data from pin 4, retrying 10 times in case of failure.
	// You may enable "boost GPIO performance" parameter, if your device is old
	// as Raspberry PI 1 (this will require root privileges). You can switch off
	// "boost GPIO performance" parameter for old devices, but it may increase
	// retry attempts. Play with this parameter.
	// temperature, humidity, _, err := dht.ReadDHTxxWithRetry(dht.DHT11, 4, false, 10)
	// if err != nil {
	// 	log.Print(err)
	// 	return
	// }

	// Print temperature and humidity

	out, err := exec.Command("airsensor.py").Output()
	if err != nil {
		log.Fatal(err)
	}

	temperature, err := strconv.ParseFloat(strings.Split(string(out[:]), " ")[0], 64)
	if err != nil {
		log.Print("Error parsing air temperature reading")
		return
	}
	log.Printf("Air Temperature = %v*C", temperature)
	airTemperature.Set(float64(temperature))

	humidity, err := strconv.ParseFloat(strings.TrimRight(strings.Split(string(out[:]), " ")[1], "\n"), 64)
	if err != nil {
		log.Print("Error parsing air humidity reading", err)
		return
	}
	log.Printf("Air Humidity = %v%%", humidity)
	airHumidity.Set(float64(humidity))
}
