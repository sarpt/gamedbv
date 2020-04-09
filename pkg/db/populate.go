package db

import (
	"github.com/jinzhu/gorm"
	"github.com/sarpt/gamedbv/pkg/db/models"
)

// Populate takes a provider of new data to be pushed to Database
func (db Database) Populate(prov PlatformProvider) error {
	db.mux.Lock()
	defer db.mux.Unlock()

	return db.handle.Transaction(func(tx *gorm.DB) error {
		if err := populatePlatform(tx, prov.Platform); err != nil {
			return err
		}

		if err := populateGames(tx, prov.Games); err != nil {
			return err
		}

		if err := populateLanguages(tx, prov.Languages); err != nil {
			return err
		}

		// manual creation of GameDescriptions instead of relying on Game.Descriptions being filled is due to horrendous performance with latter
		if err := populateDescriptions(tx, prov.Descriptions); err != nil {
			return err
		}

		return nil
	})
}

func populateLanguages(tx *gorm.DB, languages []*models.Language) error {
	for _, lang := range languages {
		identity := models.Language{Code: lang.Code}
		tx.FirstOrCreate(lang, identity)
		if tx.Error != nil {
			return tx.Error
		}
	}

	return nil
}

func populateGames(tx *gorm.DB, games []*models.Game) error {
	for _, game := range games {
		identity := models.Game{SerialNo: game.SerialNo}
		tx.Assign(game).FirstOrCreate(game, identity)
		if tx.Error != nil {
			return tx.Error
		}
	}

	return nil
}

func populateDescriptions(tx *gorm.DB, descriptions []*models.GameDescription) error {
	for _, description := range descriptions {
		identity := models.GameDescription{GameID: description.GameID, LanguageID: description.LanguageID}
		tx.Assign(description).FirstOrCreate(description, identity)
		if tx.Error != nil {
			return tx.Error
		}
	}

	return nil
}

func populatePlatform(tx *gorm.DB, platform *models.Platform) error {
	identity := models.Platform{Name: platform.Name}
	tx.Assign(platform).FirstOrCreate(platform, identity)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
