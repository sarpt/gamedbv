package models

import (
	"github.com/jinzhu/gorm"
)

// Platform inlcudes information about
type Platform struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
}
