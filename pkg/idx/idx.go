package idx

import (
	"fmt"

	"github.com/sarpt/gamedbv/pkg/index"
	"github.com/sarpt/gamedbv/pkg/parser"
	"github.com/sarpt/gamedbv/pkg/platform"
	"github.com/sarpt/gamedbv/pkg/progress"
	"github.com/sarpt/gamedbv/pkg/zip"
)

// IndexPlatform creates Index related to the platfrom
func IndexPlatform(platform platform.Variant, printer progress.Notifier) {
	printer.NextProgress(fmt.Sprintf("Unzipping platform %s", platform.String()))
	err := zip.UnzipPlatformDatabase(platform)
	if err != nil {
		printer.NextError(err)
	}

	printer.NextProgress(fmt.Sprintf("Parsing platform %s", platform.String()))
	datafile, err := parser.ParseDatabaseFile(platform)
	if err != nil {
		printer.NextError(err)
	}

	printer.NextProgress(fmt.Sprintf("Indexing platform %s", platform.String()))
	err = index.PrepareIndex(platform, datafile.Games)
	if err != nil {
		printer.NextError(err)
	}
}
