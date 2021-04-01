package idx

import (
	"errors"
	"fmt"
	"os"

	"github.com/sarpt/gamedbv/pkg/db"
	"github.com/sarpt/gamedbv/pkg/db/models"
	"github.com/sarpt/gamedbv/pkg/platform"
)

var (
	ErrDbNoAccess         = errors.New("database could not be accessed")
	ErrDbNotExist         = errors.New("database does not exist")
	ErrDbNotOpen          = errors.New("database could not be opened")
	ErrInitDbAlreadyExist = errors.New("database already exists - the initialization was not forced")
)

type InitializeDbConfig struct {
	path    string
	variant string
	force   bool
}

func (s *Server) InitializeDatabase(cfg InitializeDbConfig) error {
	var dbPath = cfg.path
	if dbPath == "" {
		dbPath = s.cfg.DbPath
	}

	var dbVariant = cfg.variant
	if dbVariant == "" {
		dbVariant = s.cfg.DbVariant
	}

	_, err := os.Stat(dbPath)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	if err == nil && !cfg.force {
		return ErrInitDbAlreadyExist
	}

	if err == nil {
		if rmErr := os.Remove(dbPath); rmErr != nil {
			return fmt.Errorf("removal of database file in path '%s' for reinitialization unsuccessful: %v", dbPath, rmErr)
		}
	}

	s.cfg.DbPath = dbPath
	s.cfg.DbVariant = dbVariant

	initalData := initialData()
	database, err := db.NewDatabase(s.dbConfig(), initalData)
	if err != nil {
		return fmt.Errorf("could not create the '%s' database in path '%s': %v", dbVariant, dbPath, err)
	}

	s.db = &database

	return nil
}

// OpenDatabase opens the database pointed to by application config.
func (s *Server) OpenDatabase() error {
	var database db.Database

	databasePath := s.cfg.DbPath
	_, err := os.Stat(databasePath)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("%w - could not access db in path %s: %v", ErrDbNoAccess, databasePath, err)
	}

	if err != nil && os.IsNotExist(err) {
		return ErrDbNotExist
	}

	database, err = db.OpenDatabase(s.dbConfig())
	if err != nil {
		s.outLog.Printf("could not open the database in path %s: %v", s.cfg.DbPath, err)
		return ErrDbNotOpen
	}

	s.db = &database

	return nil
}

func initialData() db.InitialData {
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

func (s *Server) dbConfig() db.Config {
	return db.Config{
		Path:     s.cfg.DbPath,
		Variant:  s.cfg.DbVariant,
		MaxLimit: s.cfg.DbMaxLimit,
	}
}
