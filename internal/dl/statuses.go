package dl

import (
	"fmt"

	"github.com/sarpt/gamedbv/internal/progress"
	"github.com/sarpt/gamedbv/pkg/platform"
)

const (
	process string = "dl"
	step    string = "archive-download"
)

func newArchiveFileAlreadyPresentStatus(variant platform.Variant) progress.Status {
	return progress.Status{
		Platform: variant.ID(),
		Process:  process,
		Step:     step,
		Message:  fmt.Sprintf(archiveFileAlreadyPresent, variant.Name()),
	}
}

func newDownloadingInProgressStatus(variant platform.Variant) progress.Status {
	return progress.Status{
		Platform: variant.ID(),
		Process:  process,
		Step:     step,
		Message:  fmt.Sprintf(downloadingInProgress, variant.Name()),
	}
}
