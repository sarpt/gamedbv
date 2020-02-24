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
	platform     string
	root         gametdb.Datafile
	games        map[string]*models.Game
	descriptions []*models.GameDescription
	languages    map[string]*models.Language
}

// Games returns games adapted to the format database handler accepts
func (adapt GameTDBAdapter) Games() []*models.Game {
	var games []*models.Game
	for _, game := range adapt.games {
		games = append(games, game)
	}

	return games
}

func (adapt GameTDBAdapter) GameDescriptions() []*models.GameDescription {
	return adapt.descriptions
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
		Games:        adapt.Games(),
		Descriptions: adapt.GameDescriptions(),
	}
}

// NewGameTDBAdapter returns new instance of GameTDB adapter to index and db models
func NewGameTDBAdapter(platform string, provider gametdb.ModelProvider) GameTDBAdapter {
	adapt := GameTDBAdapter{
		platform:  platform,
		root:      provider.Root(),
		games:     make(map[string]*models.Game),
		languages: make(map[string]*models.Language),
	}

	var gameDescriptions []*models.GameDescription

	for _, game := range adapt.root.Games {
		newGame := &models.Game{
			SerialNo:  game.ID,
			Region:    game.Region,
			Developer: game.Developer,
			Date:      fmt.Sprintf("%d-%d-%d", game.Date.Day, game.Date.Month, game.Date.Year),
		}
		adapt.games[game.ID] = newGame

		for _, desc := range game.Locales {
			language, ok := adapt.languages[desc.Language]
			if !ok {
				language = &models.Language{
					Code: desc.Language,
				}
				adapt.languages[language.Code] = language
			}

			description := &models.GameDescription{
				Title:    desc.Title,
				Synopsis: desc.Synopsis,
				Language: language,
				Game:     newGame,
			}

			adapt.descriptions = append(adapt.descriptions, description)
			gameDescriptions = append(gameDescriptions, description)
		}
	}

	return adapt
}
