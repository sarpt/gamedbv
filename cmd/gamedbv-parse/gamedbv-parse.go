package main

import (
	"flag"
	"fmt"

	"github.com/sarpt/gamedbv/pkg/dbunzip"
	"github.com/sarpt/gamedbv/pkg/platform"
)

func main() {
	var platformFlag platform.Variant
	var progress = make(chan string)
	defer close(progress)
	var errors = make(chan error)
	defer close(errors)

	flag.Var(&platformFlag, "platform", "platform specifies which database console variant should be parsed")
	shouldParseAllPlatforms := flag.Bool("allPlatforms", false, "When specified as true, all possible GameTDB console platforms databases will be parsed. When false, platform argument is mandatory")
	flag.Parse()

	var platformsToParse []platform.Variant
	go handleProgressMessages(progress)
	go handleErrors(errors)

	if *shouldParseAllPlatforms {
		platformsToParse = append(platformsToParse, platform.GetAllPlatforms()...)
	} else if platformFlag.IsSet() {
		platformsToParse = append(platformsToParse, platformFlag)
	} else {
		progress <- "Neither --platform nor --allPlarforms specified. One of them is mandatory."
		flag.PrintDefaults()
		return
	}

	for _, platformToParse := range platformsToParse {
		dbunzip.UnzipPlatformDatabase(platformToParse, progress, errors)
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
