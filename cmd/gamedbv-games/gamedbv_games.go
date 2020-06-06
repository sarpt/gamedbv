package main

import (
	"flag"
	"fmt"

	"github.com/sarpt/gamedbv/internal/cli"
	"github.com/sarpt/gamedbv/internal/config"
	"github.com/sarpt/gamedbv/internal/games"
	"github.com/sarpt/gamedbv/internal/progress"
	"github.com/sarpt/gamedbv/pkg/db/models"
	"github.com/sarpt/gamedbv/pkg/platform"
)

var textFlag *string
var regionFlag *string
var platformFlag *string
var jsonFlag *bool
var languageFlag *string

const (
	defaultLanguageCode string = "EN"
)

func init() {
	textFlag = flag.String("text", "", "a text to be searched for in the index")
	languageFlag = flag.String("language", defaultLanguageCode, "language code for which the description should be presented, 'EN' for english by default")
	regionFlag = flag.String("region", "", "a region of the game")
	platformFlag = flag.String("platform", "", "a platform for which the search should be executed")
	jsonFlag = flag.Bool("json", false, "when specified as true, each line of output is presented as a json object")
	flag.Parse()
}

func main() {
	var printer progress.Notifier
	if *jsonFlag {
		printer = cli.NewJSONPrinter()
	} else {
		printer = cli.NewTextPrinter()
	}

	appCfg, err := config.NewApp()
	if err != nil {
		panic(err)
	}

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

	result, err := games.Search(appCfg.Games(), params)
	if err != nil {
		panic(err)
	}

	var status progress.Status
	if *jsonFlag {
		status = progress.Status{
			Data: result,
		}
	} else {
		out := prepareOutput(result.Games, *languageFlag)
		status = progress.Status{
			Message: out,
		}
	}
	printer.NextStatus(status)
}

func prepareOutput(games []*models.Game, languageCode string) string {
	var out string

	for _, game := range games {
		out = out + fmt.Sprintf("===\n[%s]", game.SerialNo)
		for _, description := range game.Descriptions {
			if description.Language.Code == languageCode {
				out = out + fmt.Sprintf(" %s\nSynopsis: %s", description.Title, description.Synopsis)
			}
		}
		out = out + "\n"
	}

	return out
}
