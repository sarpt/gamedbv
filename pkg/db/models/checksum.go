package models

import (
	"github.com/jinzhu/gorm"
)

// Checksum provides information neccessary to check for rom validity
// A rom can have multiple calculated checksum of different variants and values
type Checksum struct {
	gorm.Model
	Variant string
	Value   string
	RomID   uint
}
