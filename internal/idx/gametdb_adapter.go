package idx

import (
	"fmt"
	"strings"

	"github.com/sarpt/gamedbv/pkg/db"
	"github.com/sarpt/gamedbv/pkg/db/models"
	"github.com/sarpt/gamedbv/pkg/gametdb"
	"github.com/sarpt/gamedbv/pkg/index"
)

const (
	unknownRegionCode string = "Other"
	gameUIDSeparator  string = ":"
	regionSeparator   string = ","
)

type gametdbModels struct {
	platform     *models.Platform
	games        map[string]*models.Game
	descriptions []*models.GameDescription
	languages    map[string]*models.Language
	gameRegions  []*models.GameRegion
	regions      map[string]*models.Region
}

type indexModels struct {
	games []index.GameSource
}

// GameTDBAdapter transforms data from GameTDB source to be usable by Index and DB packages
// This is the gluing component which should be used for transforming data provided by GameTDB source to fit any other package required
// As a result parsing, indexing and persistence packages can be decoupled using interfaces and are not dependent on the shape of GameTDB source
type GameTDBAdapter struct {
	platformID string
	root       gametdb.Datafile
	models     gametdbModels
	index      indexModels
}

// NewGameTDBAdapter returns new instance of GameTDB adapter
func NewGameTDBAdapter(platformID string, provider gametdb.ModelProvider) GameTDBAdapter {
	adapt := &GameTDBAdapter{
		platformID: platformID,
		root:       provider.Root(),
		models: gametdbModels{
			platform:  &models.Platform{UID: platformID, Indexed: true},
			games:     make(map[string]*models.Game),
			languages: make(map[string]*models.Language),
			regions:   make(map[string]*models.Region),
		},
	}

	for _, game := range adapt.root.Games {
		adapt.addGameDbModel(game)
		adapt.addGameSource(game)
	}

	return *adapt
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

func (adapt GameTDBAdapter) gameRegions() []*models.GameRegion {
	return adapt.models.gameRegions
}

func (adapt GameTDBAdapter) regions() []*models.Region {
	var regions []*models.Region
	for _, region := range adapt.models.regions {
		regions = append(regions, region)
	}

	return regions
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
		GameRegions:  adapt.gameRegions(),
		Regions:      adapt.regions(),
	}
}

func (adapt *GameTDBAdapter) addGameDbModel(source gametdb.Game) {
	newGame := &models.Game{
		UID:       adapt.generateUID(source),
		SerialNo:  source.ID,
		Developer: source.Developer,
		Date:      fmt.Sprintf("%d-%d-%d", source.Date.Day, source.Date.Month, source.Date.Year), // this needs fixing too
		Platform:  adapt.models.platform,
	}

	regionCodes := parseSourceRegion(source)
	for _, regionCode := range regionCodes {
		adapt.addGameRegion(newGame, regionCode)
	}

	for _, desc := range source.Locales {
		adapt.addDescription(newGame, desc)
	}

	adapt.models.games[newGame.UID] = newGame
}

func parseSourceRegion(source gametdb.Game) []string {
	var regionCodes []string = strings.Split(source.Region, regionSeparator)
	if len(regionCodes) == 1 && regionCodes[0] == "" {
		regionCodes = []string{unknownRegionCode}
	}

	return regionCodes
}

func (adapt *GameTDBAdapter) addGameRegion(game *models.Game, regionCode string) {
	region, ok := adapt.models.regions[regionCode]
	if !ok {
		region = adapt.addRegion(game, regionCode)
	}

	gameRegion := &models.GameRegion{
		Region: region,
		Game:   game,
	}

	adapt.models.gameRegions = append(adapt.models.gameRegions, gameRegion)
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

func (adapt *GameTDBAdapter) addRegion(game *models.Game, regionCode string) *models.Region {
	region := &models.Region{
		Code: regionCode,
	}
	adapt.models.regions[region.Code] = region

	return region
}

func (adapt *GameTDBAdapter) addLanguage(source gametdb.Locale) *models.Language {
	newLanguage := &models.Language{
		Code: source.Language,
		Name: convertLanguageCodeToName(source.Language), // temp, at the moment name will be stored in the db but it should be moved as per-presentation basis (translated name vs self name)
	}
	adapt.models.languages[newLanguage.Code] = newLanguage

	return newLanguage
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
	return fmt.Sprintf("%s%s%s", adapt.platformID, gameUIDSeparator, source.ID)
}
