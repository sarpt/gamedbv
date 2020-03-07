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

// IndexPlatform creates Index related to the platfrom
func IndexPlatform(appConf config.App, platformVariant platform.Variant, printer progress.Notifier, database db.Database) {
	platformName := platformVariant.String()

	platformConfig := appConf.Platform(platformVariant)

	printer.NextStatus(newPlatformUnzipStatus(platformName))
	err := zip.UnzipPlatformDatabase(platformConfig)
	if err != nil {
		printer.NextError(err)
		return
	}

	gametdbModelProvider := gametdb.ModelProvider{}
	printer.NextStatus(newPlatformParsingStatus(platformName))
	err = parser.ParseDatabaseFile(platformConfig, &gametdbModelProvider)
	if err != nil {
		printer.NextError(err)
		return
	}

	gametdbAdapter := NewGameTDBAdapter(platformVariant.String(), gametdbModelProvider)

	printer.NextStatus(newPlatformIndexingStatus(platformName))
	creators := map[string]index.Creator{
		"bleve": bleve.Creator{},
	}

	err = index.PrepareIndex(creators, platformConfig, gametdbAdapter.GameSources())
	if err != nil {
		printer.NextError(err)
		return
	}

	printer.NextStatus(newDatabasePopulateStatus(platformName))
	err = database.Populate(gametdbAdapter.PlatformProvider())
	if err != nil {
		printer.NextError(err)
	}
}
