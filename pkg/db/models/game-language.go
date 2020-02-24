package models

import (
	"github.com/jinzhu/gorm"
)

// GameLanguage associates game id with a language
// Game can have 1-multiple languages
type GameLanguage struct {
	gorm.Model
	Language   *Language
	LanguageID int
	Game       *Game
	GameID     int
}
