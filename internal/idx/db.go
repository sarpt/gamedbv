package idx

import (
	"os"

	"github.com/sarpt/gamedbv/internal/config"
	"github.com/sarpt/gamedbv/pkg/db"
	"github.com/sarpt/gamedbv/pkg/progress"
)

// OpenDatabase creates and opens the database pointed to by application config
func OpenDatabase(appConf config.App, printer progress.Notifier) (db.Database, error) {
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
