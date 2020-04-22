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

// WithUID filters platforms to only return platform with specified UID
func (q *PlatformsQuery) WithUID(uid string) *PlatformsQuery {
	q.handle = q.handle.Where("uid = ?", uid)

	return q
}

// FilterIndexed filters platforms so only platforms that were marked as indexed are returned
func (q *PlatformsQuery) FilterIndexed() *PlatformsQuery {
	q.handle = q.handle.Where("indexed = ?", true)

	return q
}

// First returns only first found platform
func (q *PlatformsQuery) First() models.Platform {
	platform := models.Platform{}

	q.handle.Model(&models.Platform{}).First(&platform)

	return platform
}

// Get returns platforms
func (q *PlatformsQuery) Get() []models.Platform {
	var platforms []models.Platform

	q.handle.Model(&models.Platform{}).Find(&platforms)

	return platforms
}
