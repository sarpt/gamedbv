package idx

import (
	"github.com/sarpt/gamedbv/internal/config"
	"github.com/sarpt/gamedbv/pkg/db"
	"github.com/sarpt/gamedbv/pkg/gametdb"
	"github.com/sarpt/gamedbv/pkg/index"
	"github.com/sarpt/gamedbv/pkg/index/bleve"
	"github.com/sarpt/gamedbv/pkg/parser"
	"github.com/sarpt/gamedbv/pkg/platform"
	"github.com/sarpt/gamedbv/pkg/progress"
	"github.com/sarpt/gamedbv/pkg/zip"
)

const bleveCreator string = "bleve"

// PreparePlatform unzips and parses source file, creates Index related to the platfrom and populates the database
func PreparePlatform(appConf config.App, platformVariant platform.Variant, printer progress.Notifier, database db.Database) {
	platformName := platformVariant.String()
	platformConfig := appConf.Platform(platformVariant)

	printer.NextStatus(newPlatformUnzipStatus(platformName))
	err := unzipPlatform(platformConfig)
	if err != nil {
		printer.NextError(err)
		return
	}

	printer.NextStatus(newPlatformParsingStatus(platformName))
	gametdbModelProvider, err := parsePlatformSource(platformConfig)
	if err != nil {
		printer.NextError(err)
		return
	}

	gametdbAdapter := NewGameTDBAdapter(platformVariant.String(), gametdbModelProvider)

	printer.NextStatus(newPlatformIndexingStatus(platformName))
	err = indexPlatform(platformConfig, gametdbAdapter)
	if err != nil {
		printer.NextError(err)
		return
	}

	printer.NextStatus(newDatabasePopulateStatus(platformName))
	err = populateDatabase(database, gametdbAdapter)
	if err != nil {
		printer.NextError(err)
	}
}

func unzipPlatform(platformConfig config.Platform) error {
	return zip.UnzipPlatformDatabase(platformConfig)
}

func parsePlatformSource(platformConfig config.Platform) (gametdb.ModelProvider, error) {
	gametdbModelProvider := gametdb.ModelProvider{}
	err := parser.ParseSourceFile(platformConfig, &gametdbModelProvider)

	return gametdbModelProvider, err
}

func indexPlatform(platformConfig config.Platform, gametdbAdapter GameTDBAdapter) error {
	creators := map[string]index.Creator{
		bleveCreator: bleve.BleveCreator{},
	}

	return index.PrepareIndex(creators, platformConfig, gametdbAdapter.GameSources())
}

func populateDatabase(database db.Database, gametdbAdapter GameTDBAdapter) error {
	return database.ProvidePlatformData(gametdbAdapter.PlatformProvider())
}
