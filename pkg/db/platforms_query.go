package db

import (
	"github.com/jinzhu/gorm"
	"github.com/sarpt/gamedbv/pkg/db/models"
)

// PlatformsQuery is responsible for returning the platforms information from database
type PlatformsQuery struct {
	handle *gorm.DB
}

// Get returns platforms
func (q *PlatformsQuery) Get() []models.Platform {
	var platforms []models.Platform

	q.handle.Model(&models.Platform{}).Find(&platforms)

	return platforms
}
