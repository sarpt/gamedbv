package dl

import (
	"fmt"

	"github.com/sarpt/gamedbv/pkg/progress"
)

const process string = "dl"

func newArchiveFileAlreadyPresentStatus(platform string) progress.Status {
	return progress.Status{
		Process: process,
		Step:    "archive-download",
		Message: fmt.Sprintf(archiveFileAlreadyPresent, platform),
	}
}

func newDownloadingInProgressStatus(platform string) progress.Status {
	return progress.Status{
		Process: process,
		Step:    "archive-download",
		Message: fmt.Sprintf(downloadingInProgress, platform),
	}
}
