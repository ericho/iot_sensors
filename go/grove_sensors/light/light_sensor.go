package main

import (
	"fmt"
	"time"
	"github.com/hybridgroup/gobot/platforms/intel-iot/edison"
)

func main() {
	e := edison.NewEdisonAdaptor("board")

	e.Connect()
	var d int
	for {
		time.Sleep(time.Second * 1)
		d, _ = e.AnalogRead("1")
		fmt.Println(d)
	}
}