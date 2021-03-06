package models

import (
	"github.com/jinzhu/gorm"
)

// Game contains information about a single game release
type Game struct {
	gorm.Model
	UID          string `gorm:"unique_index"`
	SerialNo     string
	GameRegions  []*GameRegion
	Languages    []*GameLanguage
	Descriptions []*GameDescription
	Developer    string
	Publisher    string
	Platform     *Platform
	PlatformID   uint
	Roms         []*Rom
	Day          uint
	Month        uint
	Year         uint
	Ratings      []*Rating
}
