package main

import (
	"flag"
	"fmt"
	"sync"

	"github.com/sarpt/gamedbv/internal/cli"
	"github.com/sarpt/gamedbv/internal/config"
	"github.com/sarpt/gamedbv/internal/idx"
	"github.com/sarpt/gamedbv/pkg/platform"
)

var platformFlag platform.Variant
var shouldParseAllPlatforms *bool

func init() {
	flag.Var(&platformFlag, "platform", "platform specifies which database console variant should be indexed")
	shouldParseAllPlatforms = flag.Bool("allPlatforms", false, "When specified as true, all possible console platforms databases will be indexed. When false, platform argument is mandatory")
	flag.Parse()
}

func main() {
	appConf, err := config.NewApp()
	if err != nil {
		panic(err)
	}

	var platformsToParse []platform.Variant

	printer := cli.NewPrinter()
	defer printer.Close()

	if *shouldParseAllPlatforms {
		platformsToParse = append(platformsToParse, platform.GetAllPlatforms()...)
	} else if platformFlag.IsSet() {
		platformsToParse = append(platformsToParse, platformFlag)
	} else {
		fmt.Println("neither --platform nor --allPlarforms specified. One of them is mandatory")
		flag.PrintDefaults()
		return
	}

	database, err := idx.OpenDatabase(appConf, printer)
	defer database.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	var wg sync.WaitGroup
	for _, platformToParse := range platformsToParse {
		wg.Add(1)

		go func(platform platform.Variant) {
			defer wg.Done()
			idx.IndexPlatform(appConf, platform, printer, database)
		}(platformToParse)
	}

	wg.Wait()
}
