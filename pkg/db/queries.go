package db

import (
	"github.com/jinzhu/gorm"
	"github.com/sarpt/gamedbv/pkg/db/models"
)

// GamesQuery is used for getting games from database
type GamesQuery struct {
	handle   *gorm.DB
	limit    int
	page     int
	maxLimit int
}

// FilterSerialNumbers filters games by matching serial numbers, if not called all games are returned
func (q *GamesQuery) FilterSerialNumbers(serialNumbers []string) *GamesQuery {
	q.handle = q.handle.Where("serial_no IN (?)", serialNumbers)
	return q
}

// FilterRegions filters games by matching their regions, if not called games from all regions are returned
func (q *GamesQuery) FilterRegions(regions []string) *GamesQuery {
	q.handle = q.handle.Where("region IN (?)", regions)
	return q
}

// Limit sets the maximum amount of games being returned
// It may exceed max limit set by the App config
func (q *GamesQuery) Limit(limit int) *GamesQuery {
	q.limit = limit
	return q
}

// Page sets the multiplication of limit. Used for paging the results
func (q *GamesQuery) Page(offset int) *GamesQuery {
	q.page = offset
	return q
}

// Get retrives games
func (q *GamesQuery) Get() ([]*models.Game, int) {
	var total int
	q.handle.Model(&models.Game{}).Count(&total)

	var games []*models.Game

	if q.limit > 0 {
		q.handle = q.handle.Limit(q.limit)
	}

	offset := getOffset(q.limit, q.page)
	if offset > 0 {
		q.handle = q.handle.Offset(offset)
	}

	q.handle.Preload("Descriptions.Language").Preload("Descriptions").Find(&games)

	return games, total
}

func getOffset(limit int, page int) int {
	return page * limit
}
