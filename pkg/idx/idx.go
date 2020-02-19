package idx

import (
	"fmt"

	"github.com/sarpt/gamedbv/pkg/config"
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
func IndexPlatform(platformVariant platform.Variant, printer progress.Notifier) {
	platformConfig := config.GetConfig(platformVariant)

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
	db, err := db.GetDatabase(platformConfig)
	defer db.Close()
	if err != nil {
		printer.NextError(err)
	}

	err = db.Populate(gametdbAdapter.PlatformProvider())
	if err != nil {
		printer.NextError(err)
	}
}
