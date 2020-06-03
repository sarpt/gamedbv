package db

import (
	"github.com/jinzhu/gorm"
	"github.com/sarpt/gamedbv/pkg/db/models"
	"github.com/sarpt/gamedbv/pkg/db/queries"
)

// transaction is an operation which should be executed on a database, preferably in a batch with other opertions
type transaction = func(db *gorm.DB) error

func createLanguagesTransaction(languages []*models.Language) transaction {
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

func createGamesTransaction(games []*models.Game) transaction {
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

func createDescriptionsTransaction(descriptions []*models.GameDescription) transaction {
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

func createUpdatePlatformsTransaction(platforms []*models.Platform) transaction {
	return func(db *gorm.DB) error {
		return execUpdatePlatformsTransaction(db, platforms)
	}
}

func execUpdatePlatformsTransaction(db *gorm.DB, platforms []*models.Platform) error {
	for _, platform := range platforms {
		q := queries.NewPlatformsQuery(db)
		identity := q.WithUID(platform.UID).First()
		if db.Error != nil {
			return db.Error
		}

		db.Model(&identity).Update(platform)
		if db.Error != nil {
			return db.Error
		}

		*platform = identity
	}

	return nil
}

func createInitPlatformsTransaction(platforms []*models.Platform) transaction {
	return func(db *gorm.DB) error {
		return execInitPlatformsTransaction(db, platforms)
	}
}

func execInitPlatformsTransaction(db *gorm.DB, platforms []*models.Platform) error {
	for _, platform := range platforms {
		identity := models.Platform{UID: platform.UID}
		db.FirstOrCreate(platform, identity)
		if db.Error != nil {
			return db.Error
		}
	}

	return nil
}

func createGameRegionsTransaction(gameRegions []*models.GameRegion) transaction {
	return func(db *gorm.DB) error {
		return execGameRegionsTransaction(db, gameRegions)
	}
}

func execGameRegionsTransaction(db *gorm.DB, gameRegions []*models.GameRegion) error {
	for _, gameRegion := range gameRegions {
		identity := models.GameRegion{GameID: gameRegion.Game.ID, RegionID: gameRegion.Region.ID}
		db.FirstOrCreate(gameRegion, identity)
		if db.Error != nil {
			return db.Error
		}
	}

	return nil
}

func createRegionsTransaction(regions []*models.Region) transaction {
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
