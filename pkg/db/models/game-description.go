package models

import (
	"github.com/jinzhu/gorm"
)

// GameDescription incorporates detailed summary of a games, along with title
// Game descriptions are localized, as such game will have 1-multiple descriptions in 1-multiple languages
type GameDescription struct {
	gorm.Model
	Language   *Language
	LanguageID int
	Title      string
	Synopsis   string
	Game       *Game
	GameID     int
}
