package main

import (
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
				go poolAds()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func poolAds() {
	out, err := exec.Command("ads1115.py").Output()
	if err != nil {
		logger.Fatal(err)
	}

	temperature, err := strconv.ParseFloat(strings.Split(string(out[:]), ",")[0], 64)
	if err != nil {
		logger.Error("Error parsing air temperature reading")
		return
	}
	logger.Info("Air Temperature = %v*C", temperature)
	airTemperature.Set(float64(temperature))

	humidity, err := strconv.ParseFloat(strings.TrimRight(strings.Split(string(out[:]), ",")[1], "\n"), 64)
	if err != nil {
		logger.Error("Error parsing air humidity reading", err)
		return
	}
	logger.Info("Air Humidity = %v%%", humidity)
	airHumidity.Set(float64(humidity))
}
