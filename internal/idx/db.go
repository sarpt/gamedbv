package idx

import (
	"os"

	"github.com/sarpt/gamedbv/internal/config"
	"github.com/sarpt/gamedbv/pkg/db"
	"github.com/sarpt/gamedbv/pkg/progress"
)

// GetDatabase creates with initialization and opens (or just opens) the database pointed to by application config
func GetDatabase(appConf config.App, printer progress.Notifier) (db.Database, error) {
	var database db.Database

	databaseConfig := appConf.Database()
	databasePath := databaseConfig.Path()
	_, err := os.Stat(databasePath)
	if err != nil && !os.IsNotExist(err) {
		return database, err
	}

	if os.IsNotExist(err) {
		printer.NextStatus(newDatabaseCreateStatus(databasePath))
		database, err = db.NewDatabase(databaseConfig)
	} else {
		printer.NextStatus(newDatabaseReuseStatus(databasePath))
		database, err = db.OpenDatabase(databaseConfig)
	}

	return database, err
}

// InitializeDatabase fills newly created database with data that is system-specific, rather than platfrom-specific.
// At the moment only all supported console platforms are being filled.
func InitializeDatabase(database db.Database) {
	// allPlatformVariants := platform.GetAllVariants()
}
