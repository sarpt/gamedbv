package queries

import (
	"github.com/jinzhu/gorm"
	"github.com/sarpt/gamedbv/pkg/db/models"
)

// LanguagesQuery is responsible for returning the languages information from database
type LanguagesQuery struct {
	handle *gorm.DB
}

// NewLanguagesQuery returns a query used for retrieving lanugages
func NewLanguagesQuery(handle *gorm.DB) *LanguagesQuery {
	return &LanguagesQuery{
		handle: handle,
	}
}

// Get returns languages
func (q *LanguagesQuery) Get() []models.Language {
	var languages []models.Language

	q.handle.Model(&models.Language{}).Find(&languages)

	return languages
}
