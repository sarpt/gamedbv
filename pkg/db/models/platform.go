package models

import (
	"github.com/jinzhu/gorm"
)

// Platform inlcudes information about
type Platform struct {
	gorm.Model
	UID     string `gorm:"unique;not null"`
	Name    string
	Indexed bool
}
