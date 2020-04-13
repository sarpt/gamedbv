package queries

import (
	"github.com/jinzhu/gorm"
	"github.com/sarpt/gamedbv/pkg/db/models"
)

// RegionsQuery is responsible for returning the regions information from database
type RegionsQuery struct {
	handle *gorm.DB
}

// NewRegionsQuery returns a query used for retrieving lanugages
func NewRegionsQuery(handle *gorm.DB) *RegionsQuery {
	return &RegionsQuery{
		handle: handle,
	}
}

// Get returns regions
func (q *RegionsQuery) Get() []models.Region {
	var regions []models.Region

	q.handle.Model(&models.Region{}).Find(&regions)

	return regions
}
