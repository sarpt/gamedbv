package idx

import (
	"os"

	"github.com/sarpt/gamedbv/pkg/db"
	"github.com/sarpt/gamedbv/pkg/db/models"
	"github.com/sarpt/gamedbv/pkg/platform"
	"github.com/sarpt/gamedbv/pkg/progress"
)

// GetDatabase creates with initialization and opens (or just opens) the database pointed to by application config
func GetDatabase(cfg db.Config, printer progress.Notifier) (db.Database, error) {
	var database db.Database

	databasePath := cfg.Path
	_, err := os.Stat(databasePath)
	if err != nil && !os.IsNotExist(err) {
		return database, err
	}

	if os.IsNotExist(err) {
		printer.NextStatus(newDatabaseCreateStatus(databasePath))
		initalData := getInitialData()
		database, err = db.NewDatabase(cfg, initalData)
	} else {
		printer.NextStatus(newDatabaseReuseStatus(databasePath))
		database, err = db.OpenDatabase(cfg)
	}

	return database, err
}

func getInitialData() db.InitialData {
	var platforms []*models.Platform

	allPlatformVariants := platform.All()
	for _, variant := range allPlatformVariants {
		platforms = append(platforms, &models.Platform{
			UID:  variant.ID(),
			Name: variant.Name(),
		})
	}

	return db.InitialData{
		Platforms: platforms,
	}
}
