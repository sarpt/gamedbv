package idx

import (
	"fmt"
	"os"

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
func IndexPlatform(appConf config.App, platformVariant platform.Variant, printer progress.Notifier) {
	platformConfig := appConf.Platform(platformVariant)
	databaseConfig := appConf.Database()
	printer.NextProgress(fmt.Sprintf("Unzipping platform %s", platformVariant.String()))
	err := zip.UnzipPlatformDatabase(platformConfig)
	if err != nil {
		printer.NextError(err)
	}

	gametdbModelProvider := gametdb.ModelProvider{}
	printer.NextProgress(fmt.Sprintf("Parsing platform %s", platformVariant.String()))
	err = parser.ParseDatabaseFile(platformConfig, &gametdbModelProvider)
	if err != nil {
		printer.NextError(err)
	}

	gametdbAdapter := NewGameTDBAdapter(platformVariant.String(), gametdbModelProvider)

	printer.NextProgress(fmt.Sprintf("Indexing platform %s", platformVariant.String()))
	creators := map[string]index.Creator{
		"bleve": bleve.Creator{},
	}

	err = index.PrepareIndex(creators, platformConfig, gametdbAdapter.GameSources())
	if err != nil {
		printer.NextError(err)
	}

	printer.NextProgress(fmt.Sprintf("Populating database for platform %s", platformVariant.String()))
	var database db.Database

	_, err = os.Stat(databaseConfig.Path())
	if err != nil && !os.IsNotExist(err) {
		printer.NextError(err)
	}

	if os.IsNotExist(err) {
		printer.NextProgress(fmt.Sprintf("Creating new database in %s", databaseConfig.Path()))
		database, err = db.NewDatabase(databaseConfig)
	} else {
		printer.NextProgress(fmt.Sprintf("Reusing database in %s", databaseConfig.Path()))
		database, err = db.OpenDatabase(databaseConfig)
	}

	defer database.Close()
	if err != nil {
		printer.NextError(err)
	}

	err = database.Populate(gametdbAdapter.PlatformProvider())
	if err != nil {
		printer.NextError(err)
	}
}
