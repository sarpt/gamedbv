package info

import (
	"github.com/sarpt/gamedbv/pkg/db"
	"github.com/sarpt/gamedbv/pkg/db/models"
)

// Platforms returtns list of platforms available in the database
func Platforms(dbConf db.Config) ([]models.Platform, error) {
	var platforms []models.Platform

	database, err := db.OpenDatabase(dbConf)
	defer database.Close()

	if err != nil {
		return platforms, err
	}

	platforms = database.NewPlatforsmQuery().Get()
	return platforms, nil
}
