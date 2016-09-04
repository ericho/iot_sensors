package main

import (
	"github.com/hybridgroup/gobot/platforms/intel-iot/edison"
	"github.com/hybridgroup/gobot/platforms/i2c"
)

func PrintLcdMessage(lcd *i2c.GroveLcdDriver, msg string) {
	lcd.Write(msg)
}

func SetAlarmColor(lcd *i2c.GroveLcdDriver) {
	lcd.SetRGB(255, 0, 0)
}

func SetWarningColor(lcd *i2c.GroveLcdDriver) {
	lcd.SetRGB(255, 255, 0)
}

func SetOkColor(lcd *i2c.GroveLcdDriver) {
	lcd.SetRGB(0, 205, 0)
}

func main() {
	e := edison.NewEdisonAdaptor("edison")
	e.Connect()
	lcd := i2c.NewGroveLcdDriver(e, "lcd")
	lcd.Start()

	SetOkColor(lcd)
	PrintLcdMessage(lcd, "Hola\nMundo..")
}
