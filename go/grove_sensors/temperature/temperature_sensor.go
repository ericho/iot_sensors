package main

import (
	"fmt"
	"time"
	"github.com/hybridgroup/gobot/platforms/intel-iot/edison"
	"github.com/hybridgroup/gobot/platforms/gpio"
)

func main() {
		e := edison.NewEdisonAdaptor("edison")
		sensor := gpio.NewGroveTemperatureSensorDriver(e, "sensor", "0")
		e.Connect()
		sensor.Start()

		for {
			fmt.Println("Current temp : ", sensor.Temperature())
			time.Sleep(time.Second * 1)
		}
}