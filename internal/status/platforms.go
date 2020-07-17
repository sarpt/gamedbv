package status

import (
	"github.com/sarpt/gamedbv/pkg/db"
	"github.com/sarpt/gamedbv/pkg/db/models"
)

// FilterIndexing specifies whether info should return all platforms, or platforms with/without index
type FilterIndexing string

const (
	// AllPlatforms instructs that all platforms in the database, regardles of their index state, should be returned
	AllPlatforms FilterIndexing = "AllPlatforms"
	// WithIndex instructs that only platforms that have been indexed should be returned
	WithIndex FilterIndexing = "WithIndex"
	// WithoutIndex instructs that only platforms that have not been indexed should be returned
	WithoutIndex FilterIndexing = "WithoutIndex"
)

// PlatformsParameters is used for specifying filters applied to platforms results
type PlatformsParameters struct {
	UID     string
	Indexed FilterIndexing
}

// Platforms returtns list of platforms available in the database
func Platforms(dbCfg db.Config, params PlatformsParameters) ([]models.Platform, error) {
	var platforms []models.Platform

	database, err := db.OpenDatabase(dbCfg)
	defer database.Close()

	if err != nil {
		return platforms, err
	}

	query := database.NewPlatformsQuery()

	if params.UID != "" {
		query.WithUID(params.UID)
	}

	if shouldFilterByIndexing(params.Indexed) {
		query.FilterIndexed(shouldFilterIndexed(params.Indexed))
	}

	platforms = query.Get()

	return platforms, nil
}

func shouldFilterByIndexing(indexed FilterIndexing) bool {
	if indexed == AllPlatforms {
		return false
	}

	return true
}

func shouldFilterIndexed(indexed FilterIndexing) bool {
	if indexed == WithIndex {
		return true
	}

	return false
}
