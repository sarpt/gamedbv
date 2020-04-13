package queries

import (
	"github.com/jinzhu/gorm"
	"github.com/sarpt/gamedbv/pkg/db/models"
)

// PlatformsQuery is responsible for returning the platforms information from database
type PlatformsQuery struct {
	handle *gorm.DB
}

// NewPlatformsQuery returns a query used for retrieving lanugages
func NewPlatformsQuery(handle *gorm.DB) *PlatformsQuery {
	return &PlatformsQuery{
		handle: handle,
	}
}

// Get returns platforms
func (q *PlatformsQuery) Get() []models.Platform {
	var platforms []models.Platform

	q.handle.Model(&models.Platform{}).Find(&platforms)

	return platforms
}
