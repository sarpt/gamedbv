package main

import (
	"flag"
	"fmt"

	"github.com/sarpt/gamedbv/pkg/cli"
	"github.com/sarpt/gamedbv/pkg/dbindex"
	"github.com/sarpt/gamedbv/pkg/dbindex/shared"
	"github.com/sarpt/gamedbv/pkg/platform"
)

var text *string
var region *string

func init() {
	text = flag.String("text", "", "A text to be searched for in the index")
	region = flag.String("region", "", "A region of the game")
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
	params := shared.SearchParameters{
		Text:    *text,
		Regions: regions,
	}

	result, err := dbindex.Search(platform.GetAllPlatforms(), params)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
