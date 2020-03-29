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
		tx.FirstOrCreate(lang, models.Language{Code: lang.Code})
		if tx.Error != nil {
			return tx.Error
		}
	}

	return nil
}

func populateGames(tx *gorm.DB, games []*models.Game) error {
	for _, game := range games {
		tx.FirstOrCreate(game, models.Game{SerialNo: game.SerialNo})
		if tx.Error != nil {
			return tx.Error
		}
	}

	return nil
}

func populateDescriptions(tx *gorm.DB, descriptions []*models.GameDescription) error {
	for _, description := range descriptions {
		tx.Create(description)
		if tx.Error != nil {
			return tx.Error
		}
	}

	return nil
}

func populatePlatform(tx *gorm.DB, platform *models.Platform) error {
	tx.FirstOrCreate(platform, models.Platform{Name: platform.Name})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
