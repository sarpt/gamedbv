package main

import (
	"flag"
	"fmt"

	"github.com/sarpt/gamedbv/pkg/dbdownload"
)

func main() {
	var platformFlag dbdownload.Platform
	var messageChannel = make(chan string)
	defer close(messageChannel)
	var errorChannel = make(chan error)
	defer close(errorChannel)

	flag.Var(&platformFlag, "platform", "platform specifies which database console variant should be fetched from the GameTDB")
	shouldDownloadAllPlatforms := flag.Bool("allPlatforms", false, "When specified as true, all possible GameTDB console platforms databases will be downloaded. When false, platform argument is mandatory")
	flag.Parse()

	go handleProgressMessages(messageChannel)
	go handleErrors(errorChannel)

	var platforms = make([]dbdownload.Platform, 0)

	if *shouldDownloadAllPlatforms {
		platforms = append(platforms, dbdownload.GetAllPlatforms()...)
	} else if platformFlag.IsSet() {
		platforms = append(platforms, platformFlag)
	} else {
		flag.PrintDefaults()
		return
	}

	for _, platform := range platforms {
		dbdownload.DownloadPlatformDatabase(platform, messageChannel, errorChannel)
	}
}

func handleProgressMessages(messageChannel <-chan string) {
	for message := range messageChannel {
		fmt.Println(message)
	}
}

func handleErrors(errorChannel <-chan error) {
	for err := range errorChannel {
		panic(err)
	}
}
