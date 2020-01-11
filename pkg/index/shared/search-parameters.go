package shared

import (
	"github.com/sarpt/gamedbv/pkg/platform"
)

// SearchParameters provide information what criteria for results are expected
type SearchParameters struct {
	Text      string
	Regions   []string
	Platforms []platform.Variant
}
