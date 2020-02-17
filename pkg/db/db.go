package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // initializes sqlite driver as per docs requirement

	"github.com/sarpt/gamedbv/pkg/db/models"
)

// Database is a wrapper around Gorm-handled SQLite3 database that exposes methods useful for GameDBV information handling
type Database struct {
	config   Config
	database *gorm.DB
}

// Close closes the underlying open db handle
func (db Database) Close() {
	db.database.Close()
}

// Populate takes a provider of new data to be pushed to Database.
// Missing information is being added, while existing information is being updated
// No data is being removed from the database
func (db Database) Populate(prov PlatformProvider) error {
	for _, game := range prov.Games {
		db.database.Create(&game)
	}

	return nil
}

// GetDatabase attempts to open the database, performing the auto-migration in the process
func GetDatabase(conf Config) (Database, error) {
	var db Database

	handle, err := gorm.Open(conf.DatabaseVariant(), conf.DatabasePath())
	if err != nil {
		return db, err
	}

	handle.AutoMigrate(
		&models.Game{},
		&models.GameLanguage{},
		&models.GameDescription{},
		&models.Rom{},
		&models.Rating{},
		&models.Checksum{},
		&models.Language{},
	)

	db = Database{
		config:   conf,
		database: handle,
	}

	return db, handle.Error
}
