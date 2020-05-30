package api

import (
	"fmt"

	"github.com/sarpt/gamedbv/internal/progress"
)

const (
	endStep string = "end"
)

// PlatformUpdateEndStatus returns progress status of finished platform update
func PlatformUpdateEndStatus(platform string) progress.Status {
	return progress.Status{
		Platform: platform,
		Step:     endStep,
		Message:  fmt.Sprintf("Update for platform %s finished", platform),
	}
}
