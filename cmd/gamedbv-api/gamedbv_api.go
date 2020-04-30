package main

import (
	"github.com/sarpt/gamedbv/internal/api"
	"github.com/sarpt/gamedbv/internal/config"
)

func main() {
	appConf, err := config.NewApp()
	if err != nil {
		panic(err)
	}

	err = api.Serve(appConf)
	if err != nil {
		panic(err)
	}
}
