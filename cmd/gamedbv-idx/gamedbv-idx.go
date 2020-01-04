package main

import (
	"flag"

	"github.com/sarpt/gamedbv/pkg/cli"
	"github.com/sarpt/gamedbv/pkg/idx"
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
	var platformsToParse []platform.Variant

	printer := cli.New()
	defer printer.Close()

	if *shouldParseAllPlatforms {
		platformsToParse = append(platformsToParse, platform.GetAllPlatforms()...)
	} else if platformFlag.IsSet() {
		platformsToParse = append(platformsToParse, platformFlag)
	} else {
		printer.NextProgress("Neither --platform nor --allPlarforms specified. One of them is mandatory.")
		flag.PrintDefaults()
		return
	}

	for _, platformToParse := range platformsToParse {
		idx.IndexPlatform(platformToParse, printer)
	}
}
