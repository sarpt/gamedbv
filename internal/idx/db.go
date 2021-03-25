package idx

import (
	"fmt"
	"os"

	"github.com/sarpt/gamedbv/pkg/db"
	"github.com/sarpt/gamedbv/pkg/db/models"
	"github.com/sarpt/gamedbv/pkg/platform"
)

// OpenDatabase creates with initialization and opens (or just opens) the database pointed to by application config.
// TODO: This should be split between opening and creating - the code should be aware of not opened database,
// but the server should be possible to be operational for GRPC to initialize/recreate database on command.
func (s *Server) OpenDatabase() error {
	var database db.Database

	databasePath := s.cfg.DbPath
	_, err := os.Stat(databasePath)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	dbConfig := db.Config{
		Path:     s.cfg.DbPath,
		Variant:  s.cfg.DbVariant,
		MaxLimit: s.cfg.DbMaxLimit,
	}

	if os.IsNotExist(err) {
		// TODO: do nothing, OpenDatabase should only open. Following in the next commit.
		initalData := getInitialData()
		database, err = db.NewDatabase(dbConfig, initalData)
	} else {
		database, err = db.OpenDatabase(dbConfig)
	}

	if err != nil {
		return fmt.Errorf("could not open the database in path %s: %v", s.cfg.DbPath, err)
	}

	s.db = &database

	return nil
}

func getInitialData() db.InitialData {
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
