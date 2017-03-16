package main

import (
	"log"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/aio"
	"gobot.io/x/gobot/platforms/firmata"
)

func main() {
	adaptor := firmata.NewAdaptor("/dev/ttyACM0")

	// Analog Inputs through firmata (arduino)
	soilSensor := aio.NewAnalogSensorDriver(adaptor, "1")
	waterPhSensor := aio.NewAnalogSensorDriver(adaptor, "3")
	work := func() {
		soilSensor.On(soilSensor.Event("data"), func(data interface{}) {
			//log.Print("soil moisture (analog voltage) ", data)
			soilHumidity.Set(float64(data.(int)))
		})
		waterPhSensor.On(waterPhSensor.Event("data"), func(data interface{}) {
			log.Print("water ph (analog voltage) ", data)
			waterPh.Set(float64(data.(int)))
		})
	}

	robot := gobot.NewRobot("sensorBot",
		[]gobot.Connection{adaptor},
		[]gobot.Device{soilSensor, waterPhSensor},
		work,
	)

	robot.Start()
}
