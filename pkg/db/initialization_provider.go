package db

import "github.com/sarpt/gamedbv/pkg/db/models"

// InitializationProvider is used for database population of initialization data (not platform dependent)
type InitializationProvider struct {
	Platforms []*models.Platform
}
