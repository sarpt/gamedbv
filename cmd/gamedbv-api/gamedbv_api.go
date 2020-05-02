package main

import (
	"flag"

	"github.com/sarpt/gamedbv/internal/api"
	"github.com/sarpt/gamedbv/internal/config"
)

var debugFlag *bool

func init() {
	debugFlag = flag.Bool("debug", false, "sets the debug mode. At the moment, its only used to ignore the origin of the websocket request")
	flag.Parse()
}

func main() {
	appCfg, err := config.NewApp()
	if err != nil {
		panic(err)
	}

	apiCfg := appCfg.API()
	apiCfg.Debug = *debugFlag

	err = api.Serve(apiCfg)
	if err != nil {
		panic(err)
	}
}
