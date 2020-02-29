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

// Games returns games matching provided serial numbers
func (db Database) Games(serialNumbers []string) []*models.Game {
	var games []*models.Game

	db.handle.Where("serial_no IN (?)", serialNumbers).Find(&games)

	for _, game := range games {
		game.Descriptions = db.GameDescriptions(*game)
	}

	return games
}

// GameDescriptions return descriptions entries for a provided game
func (db Database) GameDescriptions(game models.Game) []*models.GameDescription {
	var descriptions []*models.GameDescription

	db.handle.Model(game).Related(&descriptions)

	for _, description := range descriptions {
		description.Language = db.DescriptionLanguage(*description)
	}

	return descriptions
}

// DescriptionLanguage returns the language of provided description entry
func (db Database) DescriptionLanguage(description models.GameDescription) *models.Language {
	language := &models.Language{}

	db.handle.Model(description).Related(language)

	return language
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
