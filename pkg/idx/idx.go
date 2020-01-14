package idx

import (
	"fmt"

	"github.com/sarpt/gamedbv/pkg/index"
	"github.com/sarpt/gamedbv/pkg/index/bleve"
	"github.com/sarpt/gamedbv/pkg/parser"
	"github.com/sarpt/gamedbv/pkg/platform"
	"github.com/sarpt/gamedbv/pkg/progress"
	"github.com/sarpt/gamedbv/pkg/zip"
)

// IndexPlatform creates Index related to the platfrom
func IndexPlatform(platformVariant platform.Variant, printer progress.Notifier) {
	printer.NextProgress(fmt.Sprintf("Unzipping platform %s", platformVariant.String()))
	err := zip.UnzipPlatformDatabase(platformVariant)
	if err != nil {
		printer.NextError(err)
	}

	printer.NextProgress(fmt.Sprintf("Parsing platform %s", platformVariant.String()))
	datafile, err := parser.ParseDatabaseFile(platformVariant)
	if err != nil {
		printer.NextError(err)
	}

	printer.NextProgress(fmt.Sprintf("Indexing platform %s", platformVariant.String()))
	platformConfig := platform.GetConfig(platformVariant)
	creators := map[string]index.Creator{
		"bleve": bleve.Creator{},
	}
	err = index.PrepareIndex(creators, platformConfig, datafile.Games)

	if err != nil {
		printer.NextError(err)
	}
}
