package idx

import (
	"os"

	"github.com/sarpt/gamedbv/internal/config"
	"github.com/sarpt/gamedbv/pkg/db"
	"github.com/sarpt/gamedbv/pkg/db/models"
	"github.com/sarpt/gamedbv/pkg/platform"
	"github.com/sarpt/gamedbv/pkg/progress"
)

// GetDatabase creates with initialization and opens (or just opens) the database pointed to by application config
func GetDatabase(appConf config.App, printer progress.Notifier) (db.Database, error) {
	var database db.Database

	databaseConfig := appConf.Database()
	databasePath := databaseConfig.Path
	_, err := os.Stat(databasePath)
	if err != nil && !os.IsNotExist(err) {
		return database, err
	}

	if os.IsNotExist(err) {
		printer.NextStatus(newDatabaseCreateStatus(databasePath))
		initalData := getInitialData()
		database, err = db.NewDatabase(databaseConfig, initalData)
	} else {
		printer.NextStatus(newDatabaseReuseStatus(databasePath))
		database, err = db.OpenDatabase(databaseConfig)
	}

	return database, err
}

func getInitialData() db.InitialData {
	var platforms []*models.Platform

	allPlatformVariants := platform.All()
	for _, variant := range allPlatformVariants {
		platforms = append(platforms, &models.Platform{
			Code: variant.ID(),
			Name: variant.Name(),
		})
	}

	return db.InitialData{
		Platforms: platforms,
	}
}
