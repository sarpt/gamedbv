package index

import "github.com/sarpt/gamedbv/pkg/gametdb"

// Config provides index information and settings
type Config interface {
	GetIndexFilePath() (string, error)
	GetIndexType() string
	GetPlatform() string
}

// Searcher is used for searching in the indexes of the same type, eg. all bleve indexes could be aggregated to return batch of results
type Searcher interface {
	Search(SearchParameters) (string, error)
	AddIndex(Config) error
}

// Creator is responsible for new index creation
type Creator interface {
	CreateIndex(string, []gametdb.Game) error
}
