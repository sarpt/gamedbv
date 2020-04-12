package models

import (
	"github.com/jinzhu/gorm"
)

// Game contains information about a single game release
type Game struct {
	gorm.Model
	UID          string `gorm:"unique_index"`
	SerialNo     string
	Region       *Region
	RegionID     uint
	Languages    []*GameLanguage
	Descriptions []*GameDescription
	Developer    string
	Publisher    string
	Platform     *Platform
	PlatformID   uint
	Roms         []*Rom
	Date         string
	Ratings      []*Rating
}
