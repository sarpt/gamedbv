package models

import (
	"github.com/jinzhu/gorm"
)

// Game contains information about a single game release
type Game struct {
	gorm.Model
	SerialNo     string `gorm:"unique_index"`
	Region       string
	Languages    []*GameLanguage
	Descriptions []*GameDescription
	Developer    string
	Publisher    string
	Roms         []*Rom
	Date         string
	Ratings      []*Rating
}
