package main

import (
	"flag"
	"fmt"

	"github.com/sarpt/gamedbv/internal/cli"
	"github.com/sarpt/gamedbv/internal/cmds"
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
var pageFlag *int
var pageLimitFlag *int

const (
	defaultLanguageCode string = "EN"
)

func init() {
	textFlag = flag.String(cmds.TextFlag, "", "a text to be searched for in the index")
	languageFlag = flag.String(cmds.LanguageFlag, defaultLanguageCode, "language code for which the description should be presented, 'EN' for english by default")
	regionFlag = flag.String(cmds.RegionFlag, "", "a region of the game")
	platformFlag = flag.String(cmds.PlatformFlag, "", "a platform for which the search should be executed")
	jsonFlag = flag.Bool(cmds.JSONFlag, false, "when specified as true, each line of output is presented as a json object")
	pageFlag = flag.Int(cmds.PageFlag, 0, "when limit is set for paging, page specifies which page of results should be returned")
	pageLimitFlag = flag.Int(cmds.PageLimitFlag, 0, "limit specifies maximum number of results that are allowed to be found and reported")
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

	// todo: add possibility to pass more than one region // will do that in the next commit
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
		Page:      *pageFlag,
		PageLimit: *pageLimitFlag,
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
