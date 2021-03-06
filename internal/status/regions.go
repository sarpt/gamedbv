package status

import (
	"github.com/sarpt/gamedbv/pkg/db"
	"github.com/sarpt/gamedbv/pkg/db/models"
)

// Regions returns list of regions available in the database
func Regions(dbCfg db.Config) ([]models.Region, error) {
	var regions []models.Region

	database, err := db.OpenDatabase(dbCfg)
	defer database.Close()

	if err != nil {
		return regions, err
	}

	regions = database.NewRegionsQuery().Get()
	return regions, nil
}
