package db

import (
	"github.com/jinzhu/gorm"
	"github.com/sarpt/gamedbv/pkg/db/models"
)

// GameQuery is used for getting games from database
type GameQuery struct {
	handle *gorm.DB
}

// FilterSerialNumbers filters games by matching serial numbers, if not called all games are returned
func (q *GameQuery) FilterSerialNumbers(serialNumbers []string) *GameQuery {
	q.handle = q.handle.Where("serial_no IN (?)", serialNumbers)
	return q
}

// FilterRegions filters games by matching their regions, if not called games from all regions are returned
func (q *GameQuery) FilterRegions(regions []string) *GameQuery {
	q.handle = q.handle.Where("region IN (?)", regions)
	return q
}

// Get retrives games
func (q *GameQuery) Get() []*models.Game {
	var games []*models.Game

	q.handle.Preload("Descriptions.Language").Preload("Descriptions").Limit(100).Find(&games)

	return games
}

// GameDescriptionsQuery is used for getting descriptions of games
type GameDescriptionsQuery struct {
	handle *gorm.DB
}

// ForGame specifies for which game the description should be returned
func (q *GameDescriptionsQuery) ForGame(game models.Game) *GameDescriptionsQuery {
	q.handle = q.handle.Where("game_id = ?", game.ID)
	return q
}

// Get retrieves game description
func (q *GameDescriptionsQuery) Get() []*models.GameDescription {
	var descriptions []*models.GameDescription

	q.handle.Preload("Language").Find(&descriptions)

	return descriptions
}

// GameDescriptionLanguageQuery is used for returning language of a game description
type GameDescriptionLanguageQuery struct {
	handle      *gorm.DB
	description models.GameDescription
}

// ForGameDescription specifies for which game description the langaguage should be returned
func (q *GameDescriptionLanguageQuery) ForGameDescription(description models.GameDescription) *GameDescriptionLanguageQuery {
	q.description = description
	return q
}

// Get retrieves language
func (q *GameDescriptionLanguageQuery) Get() *models.Language {
	language := &models.Language{}

	q.handle.Model(q.description).Related(language)

	return language
}