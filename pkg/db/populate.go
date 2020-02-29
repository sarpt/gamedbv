package db

import (
	"github.com/jinzhu/gorm"
	"github.com/sarpt/gamedbv/pkg/db/models"
)

// Populate takes a provider of new data to be pushed to Database
func (db Database) Populate(prov PlatformProvider) error {
	return db.handle.Transaction(func(tx *gorm.DB) error {
		for _, lang := range prov.Languages {
			tx.FirstOrCreate(lang, models.Language{Code: lang.Code})
			if tx.Error != nil {
				return tx.Error
			}
		}

		for _, game := range prov.Games {
			tx.FirstOrCreate(game, models.Game{SerialNo: game.SerialNo})
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
