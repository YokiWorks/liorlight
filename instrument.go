package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	airTemperature = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "air_temperature_celsius",
		Help: "Current temperature of the air.",
	})
	airHumidity = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "air_humidity",
		Help: "Current humidity of the air.",
	})
	waterTemp = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "water_temperature_celsius",
		Help: "Current temperature of the water.",
	})
	waterPh = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "water_ph_voltage",
		Help: "Current water ph voltage",
	})
	soilHumidity = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "soil_humidity_voltage",
		Help: "Current humidity of the soil",
	})
	lightIntensity = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "light_intensity_lux",
		Help: "Current intensity of the light",
	})
)

func init() {
	prometheus.MustRegister(airTemperature)
	prometheus.MustRegister(airHumidity)
	prometheus.MustRegister(waterTemp)
	prometheus.MustRegister(soilHumidity)
	prometheus.MustRegister(lightIntensity)

	http.Handle("/metrics", prometheus.Handler())
	go http.ListenAndServe(":8080", nil)
}
