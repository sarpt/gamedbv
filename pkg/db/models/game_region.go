package models

import (
	"github.com/jinzhu/gorm"
)

// GameRegion associates game with a region
// Game can have 1-multiple regions
type GameRegion struct {
	gorm.Model
	Region   *Region
	RegionID uint
	Game     *Game
	GameID   uint `gorm:"index:gameregion"`
}
