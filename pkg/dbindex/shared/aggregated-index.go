package shared

import (
	"github.com/sarpt/gamedbv/pkg/platform"
)

// AggregatedIndex is used for searching in the indexes of the same type, eg. all bleve indexes could be aggregated to return batch of results
type AggregatedIndex interface {
	Search(SearchParameters) (string, error)
	Add(platform.Config) error
}
