package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // initializes sqlite driver as per docs requirement

	"github.com/sarpt/gamedbv/pkg/db/models"
)

// Database is a wrapper around Gorm-handled SQLite3 database that exposes methods useful for GameDBV information handling
type Database struct {
	config Config
	handle *gorm.DB
}

// Close closes the underlying open db handle
func (db Database) Close() {
	db.handle.Close()
}

// Populate takes a provider of new data to be pushed to Database.
// Missing information is being added, while existing information is being updated
// No data is being removed from the database
func (db Database) Populate(prov PlatformProvider) error {
	return db.handle.Transaction(func(tx *gorm.DB) error {
		for _, game := range prov.Games {
			tx.Create(game)
			if tx.Error != nil {
				return tx.Error
			}
		}

		// manual creation of GameDescription instead of relying on Game.Descriptions being filled is due to horrendous performance with latter
		for _, description := range prov.Descriptions {
			tx.Create(description)
			if tx.Error != nil {
				return tx.Error
			}
		}

		return nil
	})
}

// GamesForSerials returns games with provided serial numbers
func (db Database) GamesForSerials(serialNumbers []string) []*models.Game {
	var games []*models.Game

	db.handle.Where("serial_no IN (?)", serialNumbers).Find(&games)

	var descriptions []*models.GameDescription
	for _, game := range games {
		db.handle.Model(game).Related(&descriptions)

		language := models.Language{}
		for _, description := range descriptions {
			db.handle.Model(description).Related(&language)
			description.Language = &language
		}

		game.Descriptions = descriptions
	}

	return games
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
	}

	return db, handle.Error
}
