package models

import (
	"github.com/jinzhu/gorm"
)

// Rating includes information about variant of rating and it's value.
// Game can have 1-multiple ratings in 1-multiple languages
type Rating struct {
	gorm.Model
	Language   *Language
	LanguageID uint
	Variant    string
	Value      string
	Game       *Game
	GameID     uint
}
