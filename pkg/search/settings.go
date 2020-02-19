package search

import "github.com/sarpt/gamedbv/pkg/platform"

// Settings allow to control search execution
type Settings struct {
	Platforms []platform.Variant
	Text      string
	Regions   []string
}
