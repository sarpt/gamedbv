package main

import (
	"flag"
	"fmt"

	"github.com/sarpt/gamedbv/internal/cli"
	"github.com/sarpt/gamedbv/internal/config"
	"github.com/sarpt/gamedbv/internal/search"
	"github.com/sarpt/gamedbv/pkg/db/models"
	"github.com/sarpt/gamedbv/pkg/platform"
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
	appConf, err := config.NewApp()
	if err != nil {
		panic(err)
	}

	printer := cli.NewPrinter()
	defer printer.Close()

	// todo: add possibility to pass more than one region
	regions := []string{}
	if *region != "" {
		regions = append(regions, *region)
	}

	var platforms []platform.Variant
	if platformVariant.IsSet() {
		platforms = append(platforms, platformVariant)
	} else {
		platforms = platform.GetAllVariants()
	}

	params := search.Settings{
		Text:      *text,
		Regions:   regions,
		Platforms: platforms,
	}

	games, err := search.FindGames(appConf, params)
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
