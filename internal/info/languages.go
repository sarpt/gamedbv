package info

import (
	"github.com/sarpt/gamedbv/pkg/db"
	"github.com/sarpt/gamedbv/pkg/db/models"
)

// Languages returns list of languages available in the database
func Languages(dbCfg db.Config) ([]models.Language, error) {
	var languages []models.Language

	database, err := db.OpenDatabase(dbCfg)
	defer database.Close()

	if err != nil {
		return languages, err
	}

	languages = database.NewLanguagesQuery().Get()
	return languages, nil
}
