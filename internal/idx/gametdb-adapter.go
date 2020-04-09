package idx

import (
	"fmt"

	"github.com/sarpt/gamedbv/pkg/db"
	"github.com/sarpt/gamedbv/pkg/db/models"
	"github.com/sarpt/gamedbv/pkg/gametdb"
	"github.com/sarpt/gamedbv/pkg/index"
)

type gametdbModels struct {
	platform     *models.Platform
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
	models   gametdbModels
	index    indexModels
}

func (adapt GameTDBAdapter) games() []*models.Game {
	var games []*models.Game
	for _, game := range adapt.models.games {
		games = append(games, game)
	}

	return games
}

func (adapt GameTDBAdapter) descriptions() []*models.GameDescription {
	return adapt.models.descriptions
}

func (adapt GameTDBAdapter) languages() []*models.Language {
	var languages []*models.Language
	for _, lang := range adapt.models.languages {
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
		Platform:     adapt.models.platform,
		Games:        adapt.games(),
		Descriptions: adapt.descriptions(),
		Languages:    adapt.languages(),
	}
}

func (adapt *GameTDBAdapter) addGameDbModel(source gametdb.Game) {
	newGame := &models.Game{
		UID:       adapt.generateUID(source),
		SerialNo:  source.ID,
		Region:    source.Region,
		Developer: source.Developer,
		Date:      fmt.Sprintf("%d-%d-%d", source.Date.Day, source.Date.Month, source.Date.Year), // this needs fixing too
		Platform:  adapt.models.platform,
	}

	for _, desc := range source.Locales {
		adapt.addDescription(newGame, desc)
	}

	adapt.models.games[newGame.UID] = newGame
}

func (adapt *GameTDBAdapter) addLanguage(source gametdb.Locale) *models.Language {
	newLanguage := &models.Language{
		Code: source.Language,
		Name: convertLanguageCodeToName(source.Language), // temp, at the moment name will be stored in the db but it should be moved as per-presentation basis (translated name vs self name)
	}
	adapt.models.languages[newLanguage.Code] = newLanguage

	return newLanguage
}

func (adapt *GameTDBAdapter) addDescription(game *models.Game, source gametdb.Locale) {
	language, ok := adapt.models.languages[source.Language]
	if !ok {
		language = adapt.addLanguage(source)
	}

	description := &models.GameDescription{
		Title:    source.Title,
		Synopsis: source.Synopsis,
		Language: language,
		Game:     game,
	}

	adapt.models.descriptions = append(adapt.models.descriptions, description)
}

func (adapt *GameTDBAdapter) addGameSource(source gametdb.Game) {
	var descriptions []index.Description
	for _, locale := range source.Locales {
		descriptions = append(descriptions, index.Description{
			Synopsis: locale.Synopsis,
		})
	}

	adapt.index.games = append(adapt.index.games, index.GameSource{
		UID:          adapt.generateUID(source),
		Name:         source.Name,
		Descriptions: descriptions,
	})
}

func (adapt GameTDBAdapter) generateUID(source gametdb.Game) string {
	return fmt.Sprintf("%s:%s", adapt.platform, source.ID)
}
