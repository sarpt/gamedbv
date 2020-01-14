package main

import (
	"flag"
	"fmt"

	"github.com/sarpt/gamedbv/pkg/cli"
	"github.com/sarpt/gamedbv/pkg/index"
	"github.com/sarpt/gamedbv/pkg/platform"
	"github.com/sarpt/gamedbv/pkg/search"
)

var text *string
var region *string
var platformVariant platform.Variant

func init() {
	text = flag.String("text", "", "A text to be searched for in the index")
	region = flag.String("region", "", "A region of the game")
	flag.Var(&platformVariant, "platform", "A platform for which the search should be executed")
	flag.Parse()
}

func main() {
	printer := cli.New()
	defer printer.Close()

	// todo: add possibility to pass more than one region
	regions := []string{}
	if *region != "" {
		regions = append(regions, *region)
	}
	params := index.SearchParameters{
		Text:    *text,
		Regions: regions,
	}

	var platforms []platform.Variant
	if platformVariant.IsSet() {
		platforms = append(platforms, platformVariant)
	} else {
		platforms = platform.GetAllPlatforms()
	}
	result, err := search.Execute(platforms, params)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
