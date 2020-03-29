package db

import "github.com/sarpt/gamedbv/pkg/db/models"

// PlatformProvider is used for database population
type PlatformProvider struct {
	Platform     *models.Platform
	Games        []*models.Game
	Descriptions []*models.GameDescription
	Languages    []*models.Language
}
