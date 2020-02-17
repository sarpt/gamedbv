package idx

import (
	"fmt"

	"github.com/sarpt/gamedbv/pkg/db"
	"github.com/sarpt/gamedbv/pkg/db/models"
	"github.com/sarpt/gamedbv/pkg/gametdb"
	"github.com/sarpt/gamedbv/pkg/index"
)

// GameTDBAdapter implements interfaces neccessary for Index creation and Database population
type GameTDBAdapter struct {
	platform string
	root     gametdb.Datafile
}

// Games returns games adapted to the format database handler accepts
func (adapt GameTDBAdapter) Games() []models.Game {
	var games []models.Game
	for _, game := range adapt.root.Games {
		var descriptions []models.GameDescription
		for _, desc := range game.Locales {
			descriptions = append(descriptions, models.GameDescription{
				Title:    desc.Title,
				Synopsis: desc.Synopsis,
				Language: models.Language{
					Code: desc.Language,
				},
			})
		}

		games = append(games, models.Game{
			SerialNo:     game.ID,
			Region:       game.Region,
			Developer:    game.Developer,
			Date:         fmt.Sprintf("%d-%d-%d", game.Date.Day, game.Date.Month, game.Date.Year),
			Descriptions: descriptions,
		})
	}
	return games
}

// GameSources returns games adapted to the format indexer accepts
func (adapt GameTDBAdapter) GameSources() []index.GameSource {
	var gameSources []index.GameSource
	for _, game := range adapt.root.Games {
		gameSources = append(gameSources, index.GameSource{
			ID:     game.ID,
			Name:   game.Name,
			Region: game.Region,
		})
	}
	return gameSources
}

// PlatformProvider returns provider used to populate the database
func (adapt GameTDBAdapter) PlatformProvider() db.PlatformProvider {
	return db.PlatformProvider{
		Games: adapt.Games(),
	}
}

// NewGameTDBAdapter returns new instance of GameTDB adapter to index and db models
func NewGameTDBAdapter(platform string, provider gametdb.ModelProvider) GameTDBAdapter {
	return GameTDBAdapter{
		platform: platform,
		root:     provider.Root(),
	}
}
