package db

import (
	"github.com/jinzhu/gorm"
	"github.com/sarpt/gamedbv/pkg/db/models"
)

// LanguagesQuery is responsible for returning the languages information from database
type LanguagesQuery struct {
	handle *gorm.DB
}

// Get returns languages
func (q *LanguagesQuery) Get() []models.Language {
	var languages []models.Language

	q.handle.Model(&models.Language{}).Find(&languages)

	return languages
}
