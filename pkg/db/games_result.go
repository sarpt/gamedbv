package db

import "github.com/sarpt/gamedbv/pkg/db/models"

// GamesResult contains informations about games fetched from database
type GamesResult struct {
	Games []*models.Game
	Total int
}
