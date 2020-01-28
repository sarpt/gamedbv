package models

import (
	"github.com/jinzhu/gorm"
)

// Language includes information about short code of the language and its full name
type Language struct {
	gorm.Model
	Code string
	Name string
}
