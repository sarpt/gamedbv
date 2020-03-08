package idx

import (
	"fmt"

	"github.com/sarpt/gamedbv/pkg/db"
	"github.com/sarpt/gamedbv/pkg/db/models"
	"github.com/sarpt/gamedbv/pkg/gametdb"
	"github.com/sarpt/gamedbv/pkg/index"
)

type gametdbModels struct {
	games        map[string]*models.Game
	descriptions []*models.GameDescription
	languages    map[string]*models.Language
}

type indexModels struct {
	games []index.GameSource
}

// GameTDBAdapter transforms data from GameTDB source to be usable by Index and DB packages
// This is the gluing component which should be used for transforming data provided by GameTDB source to fit any other package required
// As a result parsing, indexing and persistence packages can be decoupled using interfaces and are not dependent on the shape of GameTDB source
type GameTDBAdapter struct {
	platform string
	root     gametdb.Datafile
	gametdb  gametdbModels
	index    indexModels
}

func (adapt GameTDBAdapter) games() []*models.Game {
	var games []*models.Game
	for _, game := range adapt.gametdb.games {
		games = append(games, game)
	}

	return games
}

func (adapt GameTDBAdapter) descriptions() []*models.GameDescription {
	return adapt.gametdb.descriptions
}

func (adapt GameTDBAdapter) languages() []*models.Language {
	var languages []*models.Language
	for _, lang := range adapt.gametdb.languages {
		languages = append(languages, lang)
	}

	return languages
}

// GameSources returns all games adapted to the format indexer accepts
func (adapt GameTDBAdapter) GameSources() []index.GameSource {
	return adapt.index.games
}

// PlatformProvider returns all data adapted to the format database accepts
func (adapt GameTDBAdapter) PlatformProvider() db.PlatformProvider {
	return db.PlatformProvider{
		Games:        adapt.games(),
		Descriptions: adapt.descriptions(),
		Languages:    adapt.languages(),
	}
}

// NewGameTDBAdapter returns new instance of GameTDB adapter
func NewGameTDBAdapter(platform string, provider gametdb.ModelProvider) GameTDBAdapter {
	adapt := &GameTDBAdapter{
		platform: platform,
		root:     provider.Root(),
		gametdb: gametdbModels{
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

func (adapt *GameTDBAdapter) addGameDbModel(source gametdb.Game) {
	newGame := &models.Game{
		SerialNo:  fmt.Sprintf("%s_%s", adapt.platform, source.ID), // turns out serial numbers are not unique across different platforms, quick workaround
		Region:    source.Region,
		Developer: source.Developer,
		Date:      fmt.Sprintf("%d-%d-%d", source.Date.Day, source.Date.Month, source.Date.Year), // this needs fixing too
	}

	for _, desc := range source.Locales {
		adapt.addDescription(newGame, desc)
	}
}

func (adapt *GameTDBAdapter) addLanguage(source gametdb.Locale) *models.Language {
	language := &models.Language{
		Code: source.Language,
	}
	adapt.gametdb.languages[language.Code] = language

	return language
}

func (adapt *GameTDBAdapter) addDescription(game *models.Game, source gametdb.Locale) {
	language, ok := adapt.gametdb.languages[source.Language]
	if !ok {
		language = adapt.addLanguage(source)
	}

	description := &models.GameDescription{
		Title:    source.Title,
		Synopsis: source.Synopsis,
		Language: language,
		Game:     game,
	}

	adapt.gametdb.descriptions = append(adapt.gametdb.descriptions, description)
}

func (adapt *GameTDBAdapter) addGameSource(source gametdb.Game) {
	var descriptions []index.Description
	for _, locale := range source.Locales {
		descriptions = append(descriptions, index.Description{
			Synopsis: locale.Synopsis,
		})
	}

	adapt.index.games = append(adapt.index.games, index.GameSource{
		ID:           fmt.Sprintf("%s_%s", adapt.platform, source.ID),
		Name:         source.Name,
		Descriptions: descriptions,
	})
}
