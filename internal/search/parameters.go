package search

import "github.com/sarpt/gamedbv/pkg/platform"

// Parameters allow to control search execution
type Parameters struct {
	Platforms []platform.Variant
	Text      string
	Regions   []string
	Page      int
	PageLimit int
}
