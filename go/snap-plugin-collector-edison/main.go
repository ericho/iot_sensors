package main

import (
	"os"

	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/ericho/iot_sensors/go/snap-plugin-collector-edison/environment"
)

func main() {
	meta := environment.Meta()
	meta.RPCType = plugin.JSONRPC
	environment.InitSensors()
	plugin.Start(meta, new(environment.Environment), os.Args[1])
}