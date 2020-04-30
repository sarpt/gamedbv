package main

import (
	"flag"
	"fmt"
	"sync"

	"github.com/sarpt/gamedbv/internal/cli"
	"github.com/sarpt/gamedbv/internal/config"
	"github.com/sarpt/gamedbv/internal/dl"
	"github.com/sarpt/gamedbv/pkg/platform"
)

var platformFlag *string
var allPlatformsFlag *bool

func init() {
	platformFlag = flag.String("platform", "", "platform specifies which console platform's database should be fetched")
	allPlatformsFlag = flag.Bool("allPlatforms", false, "When specified as true, all possible console platforms databases will be downloaded. When false, platform argument is mandatory. Takes precedence over --platfrom")
	flag.Parse()
}

func main() {
	appConf, err := config.NewApp()
	if err != nil {
		panic(err)
	}

	var platformsToDownload []platform.Variant

	printer := cli.NewPrinter()
	defer printer.Close()

	if *allPlatformsFlag {
		platformsToDownload = append(platformsToDownload, platform.All()...)
	} else if *platformFlag != "" {
		variant, err := platform.Get(*platformFlag)
		if err != nil {
			panic(err)
		}

		platformsToDownload = append(platformsToDownload, variant)
	} else {
		fmt.Println("neither --platform nor --allPlarforms specified. One of them is mandatory")
		flag.PrintDefaults()
		return
	}

	var wg sync.WaitGroup
	for _, platformToDownload := range platformsToDownload {
		wg.Add(1)

		go func(platform platform.Variant) {
			defer wg.Done()
			dl.DownloadPlatformSource(appConf.Dl(platform), platform, printer)
		}(platformToDownload)
	}

	wg.Wait()
}
