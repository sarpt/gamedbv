package main

import (
	"flag"

	"github.com/sarpt/gamedbv/pkg/cli"
	"github.com/sarpt/gamedbv/pkg/dl"
	"github.com/sarpt/gamedbv/pkg/platform"
)

var platformFlag platform.Variant
var shouldDownloadAllPlatforms *bool

func init() {
	flag.Var(&platformFlag, "platform", "platform specifies which console platform's database should be fetched")
	shouldDownloadAllPlatforms = flag.Bool("allPlatforms", false, "When specified as true, all possible console platforms databases will be downloaded. When false, platform argument is mandatory. Takes precedence over --platfrom")
	flag.Parse()
}

func main() {
	var platformsToDownload []platform.Variant

	printer := cli.New()
	defer printer.Close()

	if *shouldDownloadAllPlatforms {
		platformsToDownload = append(platformsToDownload, platform.GetAllPlatforms()...)
	} else if platformFlag.IsSet() {
		platformsToDownload = append(platformsToDownload, platformFlag)
	} else {
		printer.NextProgress("Neither --platform nor --allPlarforms specified. One of them is mandatory.")
		flag.PrintDefaults()
		return
	}

	for _, platformToDownload := range platformsToDownload {
		dl.DownloadPlatformDatabase(platformToDownload, printer)
	}
}
