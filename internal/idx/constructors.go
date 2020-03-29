package idx

import (
	"github.com/sarpt/gamedbv/pkg/db/models"
	"github.com/sarpt/gamedbv/pkg/gametdb"
)

// NewGameTDBAdapter returns new instance of GameTDB adapter
func NewGameTDBAdapter(platform string, provider gametdb.ModelProvider) GameTDBAdapter {
	adapt := &GameTDBAdapter{
		platform: platform,
		root:     provider.Root(),
		models: gametdbModels{
			platform:  &models.Platform{Name: platform},
			games:     make(map[string]*models.Game),
			languages: make(map[string]*models.Language),
		},
	}

	for _, game := range adapt.root.Games {
		adapt.addGameDbModel(game)
		adapt.addGameSource(game)
	}

	return *adapt
}
