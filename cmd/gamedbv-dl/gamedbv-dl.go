package main

import (
	"flag"
	"fmt"

	"github.com/sarpt/gamedbv/pkg/dbdl"
	"github.com/sarpt/gamedbv/pkg/platform"
)

func main() {
	var platformFlag platform.Variant
	var progress = make(chan string)
	defer close(progress)
	var errors = make(chan error)
	defer close(errors)

	flag.Var(&platformFlag, "platform", "platform specifies which console platform's database should be fetched from the GameTDB")
	shouldDownloadAllPlatforms := flag.Bool("allPlatforms", false, "When specified as true, all possible GameTDB console platforms databases will be downloaded. When false, platform argument is mandatory. Takes precedence over --platfrom")
	flag.Parse()

	go handleProgressMessages(progress)
	go handleErrors(errors)

	var platformsToDownload []platform.Variant

	if *shouldDownloadAllPlatforms {
		platformsToDownload = append(platformsToDownload, platform.GetAllPlatforms()...)
	} else if platformFlag.IsSet() {
		platformsToDownload = append(platformsToDownload, platformFlag)
	} else {
		progress <- "Neither --platform nor --allPlarforms specified. One of them is mandatory."
		flag.PrintDefaults()
		return
	}

	for _, platformToDownload := range platformsToDownload {
		dbdl.DownloadPlatformDatabase(platformToDownload, progress, errors)
	}
}

func handleProgressMessages(progress <-chan string) {
	for message := range progress {
		fmt.Println(message)
	}
}

func handleErrors(errors <-chan error) {
	for err := range errors {
		panic(err)
	}
}
