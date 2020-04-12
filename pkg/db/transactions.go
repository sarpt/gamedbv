package db

import (
	"github.com/jinzhu/gorm"
	"github.com/sarpt/gamedbv/pkg/db/models"
)

func createLanguagesTransaction(languages []*models.Language) Transaction {
	return func(db *gorm.DB) error {
		return execLanguagesTransaction(db, languages)
	}
}

func execLanguagesTransaction(db *gorm.DB, languages []*models.Language) error {
	for _, lang := range languages {
		identity := models.Language{Code: lang.Code}
		db.FirstOrCreate(lang, identity)
		if db.Error != nil {
			return db.Error
		}
	}

	return nil
}

func createGamesTransaction(games []*models.Game) Transaction {
	return func(db *gorm.DB) error {
		return execGamesTransaction(db, games)
	}
}

func execGamesTransaction(db *gorm.DB, games []*models.Game) error {
	for _, game := range games {
		identity := models.Game{SerialNo: game.SerialNo}
		db.FirstOrCreate(game, identity)
		if db.Error != nil {
			return db.Error
		}
	}

	return nil
}

func createDescriptionsTransaction(descriptions []*models.GameDescription) Transaction {
	return func(db *gorm.DB) error {
		return execDescriptionsTransaction(db, descriptions)
	}
}

func execDescriptionsTransaction(db *gorm.DB, descriptions []*models.GameDescription) error {
	for _, description := range descriptions {
		identity := models.GameDescription{GameID: description.Game.ID, LanguageID: description.Language.ID}
		db.FirstOrCreate(description, identity)
		if db.Error != nil {
			return db.Error
		}
	}

	return nil
}

func createPlatformsTransaction(platform []*models.Platform) Transaction {
	return func(db *gorm.DB) error {
		return execPlatformTransaction(db, platform)
	}
}

func execPlatformTransaction(db *gorm.DB, platforms []*models.Platform) error {
	for _, platforms := range platforms {
		identity := models.Platform{Name: platforms.Name}
		db.FirstOrCreate(platforms, identity)
		if db.Error != nil {
			return db.Error
		}
	}

	return nil
}

func createRegionsTransaction(regions []*models.Region) Transaction {
	return func(db *gorm.DB) error {
		return execRegionsTransaction(db, regions)
	}
}

func execRegionsTransaction(db *gorm.DB, regions []*models.Region) error {
	for _, region := range regions {
		identity := models.Region{Code: region.Code}
		db.FirstOrCreate(region, identity)
		if db.Error != nil {
			return db.Error
		}
	}

	return nil
}
