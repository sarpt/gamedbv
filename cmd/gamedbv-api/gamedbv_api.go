package main

import (
	"github.com/sarpt/gamedbv/internal/api"
	"github.com/sarpt/gamedbv/internal/config"
)

func main() {
	appCfg, err := config.NewApp()
	if err != nil {
		panic(err)
	}

	err = api.Serve(appCfg.API())
	if err != nil {
		panic(err)
	}
}
