package main

import (
	"flag"
	"fmt"

	"github.com/sarpt/gamedbv/internal/cli"
	"github.com/sarpt/gamedbv/internal/config"
	"github.com/sarpt/gamedbv/internal/games"
	"github.com/sarpt/gamedbv/pkg/db/models"
	"github.com/sarpt/gamedbv/pkg/platform"
)

var textFlag *string
var regionFlag *string
var platformFlag *string

func init() {
	textFlag = flag.String("text", "", "A text to be searched for in the index")
	regionFlag = flag.String("region", "", "A region of the game")
	platformFlag = flag.String("platform", "", "A platform for which the search should be executed")
	flag.Parse()
}

func main() {

	appCfg, err := config.NewApp()
	if err != nil {
		panic(err)
	}

	printer := cli.NewTextPrinter()
	defer printer.Close()

	// todo: add possibility to pass more than one region
	regions := []string{}
	if *regionFlag != "" {
		regions = append(regions, *regionFlag)
	}

	var platforms []platform.Variant
	if *platformFlag != "" {
		variant, err := platform.Get(*platformFlag)
		if err != nil {
			panic(err)
		}

		platforms = append(platforms, variant)
	} else {
		platforms = platform.All()
	}

	params := games.SearchParameters{
		Text:      *textFlag,
		Regions:   regions,
		Platforms: platforms,
	}

	games, err := games.Search(appCfg.Games(), params)
	if err != nil {
		panic(err)
	}

	out := prepareOutput(games.Games)
	fmt.Println(out)
}

func prepareOutput(games []*models.Game) string {
	var out string

	for _, game := range games {
		for _, description := range game.Descriptions {
			if description.Language.Code == "EN" {
				out = out + fmt.Sprintf("===\n[%s] %s\nSynopsis: %s\n", game.SerialNo, description.Title, description.Synopsis)
			}
		}
	}

	return out
}
