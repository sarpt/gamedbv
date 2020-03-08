package db

import (
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // initializes sqlite driver as per docs requirement

	"github.com/sarpt/gamedbv/pkg/db/models"
)

// Database is a wrapper around Gorm-handled SQLite3 database that exposes methods useful for GameDBV information handling
type Database struct {
	config Config
	handle *gorm.DB
	mux    *sync.Mutex
}

// Close closes the underlying open db handle
func (db Database) Close() {
	if db.handle != nil {
		db.handle.Close()
	}
}

// NewGameQuery returns an object used retrieving games
func (db Database) NewGameQuery() *GameQuery {
	return &GameQuery{
		handle:                db.handle.New(),
		gameDescriptionsQuery: db.NewGameDescriptionsQuery(),
	}
}

// NewGameDescriptionsQuery returns an object used for retrieving games descriptions
func (db Database) NewGameDescriptionsQuery() *GameDescriptionsQuery {
	return &GameDescriptionsQuery{
		handle:                   db.handle.New(),
		descriptionLanguageQuery: db.NewGameDescriptionLanguageQuery(),
	}
}

// NewGameDescriptionLanguageQuery returns an object used for retriving language of a query
func (db Database) NewGameDescriptionLanguageQuery() *GameDescriptionLanguageQuery {
	return &GameDescriptionLanguageQuery{
		handle: db.handle.New(),
	}
}

// NewDatabase attempts to open the database, performing the auto-migration in the process
func NewDatabase(conf Config) (Database, error) {
	db, err := OpenDatabase(conf)
	if err != nil {
		return db, err
	}

	db.handle.AutoMigrate(
		&models.Game{},
		&models.GameLanguage{},
		&models.GameDescription{},
		&models.Rom{},
		&models.Rating{},
		&models.Checksum{},
		&models.Language{},
	)

	return db, err
}

// OpenDatabase attempts to open the database
func OpenDatabase(conf Config) (Database, error) {
	var db Database

	handle, err := gorm.Open(conf.Variant(), conf.Path())
	if err != nil {
		return db, err
	}

	db = Database{
		config: conf,
		handle: handle,
		mux:    &sync.Mutex{},
	}

	return db, handle.Error
}
