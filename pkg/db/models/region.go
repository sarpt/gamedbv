package models

import "github.com/jinzhu/gorm"

// Region contains information about place/region of game release
type Region struct {
	gorm.Model
	Code string `gorm:"unique_index"`
}
