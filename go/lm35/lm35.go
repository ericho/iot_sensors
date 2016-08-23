package main

import (
	"fmt"
	"github.com/hybridgroup/gobot/platforms/intel-iot/edison"
	"time"
)

const (
	LM35_RESOLUTION   = 0.01
	EDISON_ADCBITS    = 1024
	REFERENCE_VOLTAGE = 5
)

func computeTemp(s int) float64 {
	r := float64(s)
	return r * REFERENCE_VOLTAGE / EDISON_ADCBITS / LM35_RESOLUTION
}

func main() {
	e := edison.NewEdisonAdaptor("edison")
	e.Connect()
	var v int
	for {
		v, _ = e.AnalogRead("0")
		fmt.Printf("Temp: %f degrees.\n", computeTemp(v))
		time.Sleep(1000 * time.Millisecond)
	}
}
