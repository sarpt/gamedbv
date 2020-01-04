package idx

import (
	"fmt"

	"github.com/sarpt/gamedbv/pkg/dbindex"
	"github.com/sarpt/gamedbv/pkg/dbparse"
	"github.com/sarpt/gamedbv/pkg/dbunzip"
	"github.com/sarpt/gamedbv/pkg/platform"
	"github.com/sarpt/gamedbv/pkg/progress"
)

// IndexPlatform creates Index related to the platfrom
func IndexPlatform(platform platform.Variant, printer progress.Notifier) {
	printer.NextProgress(fmt.Sprintf("Unzipping platform %s", platform.String()))
	err := dbunzip.UnzipPlatformDatabase(platform)
	if err != nil {
		printer.NextError(err)
	}

	printer.NextProgress(fmt.Sprintf("Parsing platform %s", platform.String()))
	datafile, err := dbparse.ParseDatabaseFile(platform)
	if err != nil {
		printer.NextError(err)
	}

	printer.NextProgress(fmt.Sprintf("Indexing platform %s", platform.String()))
	err = dbindex.PrepareIndex(platform, datafile.Games)
	if err != nil {
		printer.NextError(err)
	}
}
