package db

import "github.com/sarpt/gamedbv/pkg/db/models"

// PlatformProvider is used for database population
type PlatformProvider struct {
	Games     []models.Game
	Languages []models.Language
}
