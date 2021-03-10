package games

import "github.com/sarpt/gamedbv/pkg/platform"

// SearchParameters allow to control search execution
type SearchParameters struct {
	Platforms []platform.Variant
	Text      string
	Regions   []string
	Page      int
	PageLimit int
}
