package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // initializes sqlite driver as per docs requirement

	"github.com/sarpt/gamedbv/pkg/db/models"
)

// Database is a wrapper around Gorm-handled SQLite3 database that exposes methods useful for GameDBV information handling
type Database struct {
	config Config
	db     *gorm.DB
}

// Close closes the underlying open db handle
func (db Database) Close() {
	db.Close()
}

// GetDatabase attempts to open the database, performing the auto-migration in the process
// Remember to close the database
func GetDatabase(conf Config) (Database, error) {
	var db Database

	handle, err := gorm.Open("sqlite", conf.Path())
	if err != nil {
		return db, err
	}

	handle.AutoMigrate(&models.Game{})

	db = Database{
		config: conf,
		db:     handle,
	}

	return db, nil
}
