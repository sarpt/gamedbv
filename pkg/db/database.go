package db

import (
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // initializes sqlite driver as per docs requirement

	"github.com/sarpt/gamedbv/pkg/db/models"
)

// Transaction is an operation which should be executed on a database, preferably in a batch with other opertions
type Transaction = func(db *gorm.DB) error

// Database is a wrapper around Gorm-handled SQLite3 database that exposes methods useful for GameDBV information handling
type Database struct {
	config Config
	handle *gorm.DB
	mux    *sync.Mutex
}

// NewDatabase attempts to open the database, performing the auto-migration in the process
func NewDatabase(conf Config) (Database, error) {
	db, err := OpenDatabase(conf)
	if err != nil {
		return db, err
	}

	db.handle.AutoMigrate(
		&models.Platform{},
		&models.Game{},
		&models.GameLanguage{},
		&models.GameDescription{},
		&models.Rom{},
		&models.Rating{},
		&models.Checksum{},
		&models.Language{},
		&models.Region{},
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

// Close closes the underlying open db handle
func (db Database) Close() {
	if db.handle != nil {
		db.handle.Close()
	}
}

// NewGamesQuery returns a query used for retrieving games
func (db Database) NewGamesQuery() *GamesQuery {
	return &GamesQuery{
		handle:   db.handle.New(),
		maxLimit: db.config.MaxLimit(),
	}
}

// NewLanguagesQuery returns a query used for retrieving lanugages
func (db Database) NewLanguagesQuery() *LanguagesQuery {
	return &LanguagesQuery{
		handle: db.handle.New(),
	}
}

// NewPlatforsmQuery returns a query used for retrieving lanugages
func (db Database) NewPlatforsmQuery() *PlatformsQuery {
	return &PlatformsQuery{
		handle: db.handle.New(),
	}
}

// ExecuteTransactions locks the database in order to execute batch of operations
func (db Database) ExecuteTransactions(transactions []Transaction) error {
	db.mux.Lock()
	defer db.mux.Unlock()

	return db.handle.Transaction(func(tx *gorm.DB) error {
		for _, transaction := range transactions {
			if err := transaction(tx); err != nil {
				return tx.Error
			}
		}

		return nil
	})
}

// ProvidePlatformData takes a provider of platform's new data to be pushed to Database
func (db Database) ProvidePlatformData(provider PlatformProvider) error {
	transactions := []Transaction{
		createPlatformsTransaction([]*models.Platform{provider.Platform}),
		createRegionsTransaction(provider.Regions),
		createGamesTransaction(provider.Games),
		createLanguagesTransaction(provider.Languages),
		createDescriptionsTransaction(provider.Descriptions),
	}

	return db.ExecuteTransactions(transactions)
}

// ProvideInitialData fills data needed for application work (independent from platforms etc.)
func (db Database) ProvideInitialData(provider InitializationProvider) error {
	transactions := []Transaction{
		createPlatformsTransaction(provider.Platforms),
	}

	return db.ExecuteTransactions(transactions)
}
