package idx

import (
	"github.com/sarpt/gamedbv/pkg/db/models"
	"github.com/sarpt/gamedbv/pkg/gametdb"
	"github.com/sarpt/gamedbv/pkg/index"
)

// GameTdbAdapter implements interfaces neccessary for Index creation and Database population
type GameTdbAdapter struct {
	platform string
	games    []gametdb.Game
}

// Games returns games adapted to the format database handler accepts
func (pp GameTdbAdapter) Games() []models.Game {
	var games []models.Game
	for _, game := range pp.games {
		games = append(games, models.Game{
			SerialNo: game.ID,
			Region:   game.Region,
		})
	}
	return games
}

// GameSources returns games adapted to the format indexer accepts
func (pp GameTdbAdapter) GameSources() []index.GameSource {
	var gameSources []index.GameSource
	for _, game := range pp.games {
		gameSources = append(gameSources, index.GameSource{
			ID:     game.ID,
			Name:   game.Name,
			Region: game.Region,
		})
	}
	return gameSources
}

// NewGameTdbAdapter returns new instance of GameTDB adapter to index and db models
func NewGameTdbAdapter(platform string, provider gametdb.ModelProvider) GameTdbAdapter {
	return GameTdbAdapter{
		platform: platform,
		games:    provider.Games(),
	}
}
