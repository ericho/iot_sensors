package environment

import (
	"github.com/hybridgroup/gobot/platforms/intel-iot/edison"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/hybridgroup/gobot/platforms/i2c"
)

const (
	TempSensorPort = "0"
	LightSensorPort = "1"
)

var board *edison.EdisonAdaptor
var tempSensor *gpio.GroveTemperatureSensorDriver
var lcd *i2c.GroveLcdDriver

func InitSensors() {
	board = edison.NewEdisonAdaptor("board")
	tempSensor = gpio.NewGroveTemperatureSensorDriver(board, "sensor", TempSensorPort)
	lcd = i2c.NewGroveLcdDriver(board, "lcd")
	board.Connect()

	tempSensor.Start()
	lcd.Start()
}

func Temperature() float64 {
	PrintLcdMessage("Temp")
	return tempSensor.Temperature()
}

func Ligth() int {
	value, _ := board.AnalogRead(LightSensorPort)
	return value
}

func PrintLcdMessage(msg string) {
	lcd.Write(msg)
}

func SetAlarmColor() {
	lcd.SetRGB(255, 0, 0)
}

func SetWarningColor() {
	lcd.SetRGB(255, 255, 0)
}

func SetOkColor() {
	lcd.SetRGB(0, 205, 0)
}

