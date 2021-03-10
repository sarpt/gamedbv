package main

import (
	"flag"
	"os"

	"github.com/sarpt/gamedbv/internal/api"
	"github.com/sarpt/gamedbv/internal/cmds"
	"github.com/sarpt/gamedbv/internal/config"
)

var debugFlag *bool
var interfaceFlag *string

func init() {
	debugFlag = flag.Bool(cmds.DebugFlag, false, "sets the debug mode. At the moment, its only used to ignore the origin of the websocket request")
	interfaceFlag = flag.String(cmds.InterfaceFlag, "", "when set, address on which the API is served is taken from the interface matching the name, selecting the first non-loopback ipv4 address. When arguemnt is set and interface is not found or it does not have non-loopback ipv4 address, programs falls back to localhost")
	flag.Parse()
}

func main() {
	projectCfg, err := config.Create()
	if err != nil {
		panic(err)
	}

	apiCfg := projectCfg.API
	apiCfg.Debug = *debugFlag
	apiCfg.NetInterface = *interfaceFlag

	server := api.NewServer(apiCfg)
	err = server.Serve(os.Stdout)
	if err != nil {
		panic(err)
	}
}
