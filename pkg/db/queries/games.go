package queries

import (
	"github.com/jinzhu/gorm"
	"github.com/sarpt/gamedbv/pkg/db/models"
)

// GamesResult contains informations about games fetched from database
type GamesResult struct {
	Games []*models.Game
	Total int
}

// GamesQuery is used for getting games from database
type GamesQuery struct {
	handle   *gorm.DB
	limit    int
	page     int
	maxLimit int
	regions  []string
}

// NewGamesQuery return the new query used for retrieving games
func NewGamesQuery(handle *gorm.DB, maxLimit int) *GamesQuery {
	return &GamesQuery{
		handle:   handle,
		maxLimit: maxLimit,
	}
}

// FilterUIDs filters games by matching UID (platform:serial_no), if not called all games are returned
func (q *GamesQuery) FilterUIDs(serialNumbers []string) *GamesQuery {
	q.handle = q.handle.Where("games.uid IN (?)", serialNumbers)
	return q
}

// FilterRegions filters games by matching their regions, if not called games from all regions are returned
func (q *GamesQuery) FilterRegions(regions []string) *GamesQuery {
	q.handle = q.handle.Joins("left join game_regions on game_regions.game_id = games.id").Joins("left join regions on regions.id = game_regions.region_id").Where("regions.code IN (?)", regions)
	return q
}

// FilterPlatforms filters games by matching platforms uids, if not called games on all platforms are returned
func (q *GamesQuery) FilterPlatforms(platforms []string) *GamesQuery {
	q.handle = q.handle.Joins("left join platforms on platforms.id = games.platform_id").Where("platforms.uid IN (?)", platforms)
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

// Get retrives games and their total count
func (q *GamesQuery) Get() GamesResult {
	var total int
	q.handle.Model(&models.Game{}).Count(&total)

	var games []*models.Game

	var limit = q.limit
	if limit <= 0 {
		limit = total
	}

	neededOffsets := offsets(limit, q.maxLimit, q.page)

	for page, offset := range neededOffsets {
		var gamesForPage []*models.Game
		pageLimit := limitForPage(page, len(neededOffsets), limit, q.maxLimit)

		q.handle.Limit(pageLimit).Offset(offset).Preload("GameRegions.Region").Preload("GameRegions").Preload("Descriptions.Language").Preload("Descriptions").Preload("Platform").Find(&gamesForPage)

		games = append(games, gamesForPage...)
	}

	return GamesResult{
		Games: games,
		Total: total,
	}
}
