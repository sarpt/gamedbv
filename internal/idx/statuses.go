package idx

import (
	"fmt"

	"github.com/sarpt/gamedbv/pkg/progress"
)

const process string = "idx"

func newPlatformUnzipStatus(platform string) progress.Status {
	return progress.Status{
		Process: process,
		Step:    "unzip",
		Message: fmt.Sprintf("Unzipping platform %s", platform),
	}
}

func newPlatformParsingStatus(platform string) progress.Status {
	return progress.Status{
		Process: process,
		Step:    "parse",
		Message: fmt.Sprintf("Parsing platform %s", platform),
	}
}

func newPlatformIndexingStatus(platform string) progress.Status {
	return progress.Status{
		Process: process,
		Step:    "index",
		Message: fmt.Sprintf("Indexing platform %s", platform),
	}
}

func newDatabasePopulateStatus(platform string) progress.Status {
	return progress.Status{
		Process: process,
		Step:    "database-populate",
		Message: fmt.Sprintf("Populating database for platform %s", platform),
	}
}

func newDatabaseCreateStatus(path string) progress.Status {
	return progress.Status{
		Process: process,
		Step:    "database-create",
		Message: fmt.Sprintf("Creating new database in %s", path),
	}
}

func newDatabaseReuseStatus(path string) progress.Status {
	return progress.Status{
		Process: process,
		Step:    "database-reuse",
		Message: fmt.Sprintf("Reusing database in %s", path),
	}
}
