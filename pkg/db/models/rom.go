package models

import (
	"github.com/jinzhu/gorm"
)

// Rom includes information about known roms of a game
// A game can have 0-multiple roms
type Rom struct {
	gorm.Model
	Version   string
	Name      string
	Size      string
	Checksums []Checksum
	GameID    int
}
