package main

import (
	"flag"
	"sync"

	"github.com/sarpt/gamedbv/internal/cli"
	"github.com/sarpt/gamedbv/internal/cmds"
	"github.com/sarpt/gamedbv/internal/config"
	"github.com/sarpt/gamedbv/internal/dl"
	"github.com/sarpt/gamedbv/internal/progress"
	"github.com/sarpt/gamedbv/pkg/platform"
)

var jsonFlag *bool
var platformsFlag *cmds.MultipleFlag = &cmds.MultipleFlag{}

func init() {
	flag.Var(platformsFlag, cmds.PlatformFlag, "platform specifies which console platform's database should be fetched")
	jsonFlag = flag.Bool(cmds.JSONFlag, false, "when specified as true, each line of output is presented as a json object")
	flag.Parse()
}

func main() {
	appCfg, err := config.NewApp()
	if err != nil {
		panic(err)
	}

	var platformsToDownload []platform.Variant

	var printer progress.Notifier
	if *jsonFlag {
		printer = cli.NewJSONPrinter()
	} else {
		printer = cli.NewTextPrinter()
	}

	if len(platformsFlag.Values()) == 0 {
		platformsToDownload = append(platformsToDownload, platform.All()...)
	} else {
		for _, val := range platformsFlag.Values() {
			variant, err := platform.Get(val)
			if err != nil {
				panic(err)
			}

			platformsToDownload = append(platformsToDownload, variant)
		}
	}

	var wg sync.WaitGroup
	for _, platformToDownload := range platformsToDownload {
		wg.Add(1)

		go func(platform platform.Variant) {
			defer wg.Done()
			dl.DownloadPlatformSource(appCfg.Dl(platform), platform, printer)
		}(platformToDownload)
	}

	wg.Wait()
}
