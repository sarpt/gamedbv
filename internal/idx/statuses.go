package idx

import (
	"fmt"

	"github.com/sarpt/gamedbv/internal/progress"
	"github.com/sarpt/gamedbv/pkg/platform"
)

const (
	process      string = "idx"
	unzipStep    string = "unzip"
	parseStep    string = "parse"
	indexStep    string = "index"
	populateStep string = "database-populate"
	createStep   string = "database-create"
	reuseStep    string = "database-reuse"
)

func newPlatformUnzipStatus(variant platform.Variant) progress.Status {
	return progress.Status{
		Platform: variant.ID(),
		Process:  process,
		Step:     unzipStep,
		Message:  fmt.Sprintf("Unzipping platform %s", variant.Name()),
	}
}

func newPlatformParsingStatus(variant platform.Variant) progress.Status {
	return progress.Status{
		Platform: variant.ID(),
		Process:  process,
		Step:     parseStep,
		Message:  fmt.Sprintf("Parsing platform %s", variant.Name()),
	}
}

func newPlatformIndexingStatus(variant platform.Variant) progress.Status {
	return progress.Status{
		Platform: variant.ID(),
		Process:  process,
		Step:     indexStep,
		Message:  fmt.Sprintf("Indexing platform %s", variant.Name()),
	}
}

func newDatabasePopulateStatus(variant platform.Variant) progress.Status {
	return progress.Status{
		Platform: variant.ID(),
		Process:  process,
		Step:     populateStep,
		Message:  fmt.Sprintf("Populating database for platform %s", variant.Name()),
	}
}

func newDatabaseCreateStatus(path string) progress.Status {
	return progress.Status{
		Process: process,
		Step:    createStep,
		Message: fmt.Sprintf("Creating new database in %s", path),
	}
}

func newDatabaseReuseStatus(path string) progress.Status {
	return progress.Status{
		Process: process,
		Step:    reuseStep,
		Message: fmt.Sprintf("Reusing database in %s", path),
	}
}
