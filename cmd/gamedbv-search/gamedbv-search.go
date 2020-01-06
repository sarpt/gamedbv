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

func init() {
	text = flag.String("text", "", "A text to be searched for in the index")
	flag.Parse()
}

func main() {
	printer := cli.New()
	defer printer.Close()

	var regions []string
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
