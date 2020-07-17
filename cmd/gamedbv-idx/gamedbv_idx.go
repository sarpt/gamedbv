package main

import (
	"flag"
	"fmt"
	"sync"

	"github.com/sarpt/goutils/pkg/listflag"

	"github.com/sarpt/gamedbv/internal/cli"
	"github.com/sarpt/gamedbv/internal/cmds"
	"github.com/sarpt/gamedbv/internal/config"
	"github.com/sarpt/gamedbv/internal/idx"
	"github.com/sarpt/gamedbv/internal/progress"
	"github.com/sarpt/gamedbv/pkg/platform"
)

var jsonFlag *bool
var platformFlags *listflag.StringList

func init() {
	platformFlags = listflag.NewStringList([]string{})

	flag.Var(platformFlags, cmds.PlatformFlag, "platform specifies which console platform's database should be fetched")
	jsonFlag = flag.Bool(cmds.JSONFlag, false, "when specified as true, each line of output is presented as a json object")
	flag.Parse()
}

func main() {
	appCfg, err := config.NewApp()
	if err != nil {
		panic(err)
	}

	var platformsToParse []platform.Variant

	var printer progress.Notifier
	if *jsonFlag {
		printer = cli.NewJSONPrinter()
	} else {
		printer = cli.NewTextPrinter()
	}

	if len(platformFlags.Values()) == 0 {
		platformsToParse = append(platformsToParse, platform.All()...)
	} else {
		for _, val := range platformFlags.Values() {
			variant, err := platform.Get(val)
			if err != nil {
				panic(err)
			}

			platformsToParse = append(platformsToParse, variant)
		}
	}

	database, err := idx.GetDatabase(appCfg.Database(), printer)
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
			idx.PreparePlatform(appCfg.Idx(platform), platform, printer, database)
		}(platformToParse)
	}

	wg.Wait()
}
