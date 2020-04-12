package main

import (
	"github.com/sarpt/gamedbv/internal/config"
	"github.com/sarpt/gamedbv/internal/serv"
)

func main() {
	appConf, err := config.NewApp()
	if err != nil {
		panic(err)
	}

	err = serv.Serve(appConf)
	if err != nil {
		panic(err)
	}
}
