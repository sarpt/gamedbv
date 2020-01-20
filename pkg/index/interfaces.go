package index

// Config provides index information and settings
type Config interface {
	GetIndexFilePath() (string, error)
	GetIndexType() string
	GetPlatform() string
	GetDocType() string
}

// Searcher is used for searching in the indexes of the same type, eg. all bleve indexes could be aggregated to return batch of results
type Searcher interface {
	Search(SearchParameters) (Result, error)
	AddIndex(Config) error
}

// Creator is responsible for new index creation
type Creator interface {
	CreateIndex(string, string, []GameSource) error
}

// GameSource implements methods returning information about Game neccessary for index construction
type GameSource interface {
	GetID() string
	GetName() string
	GetRegion() string
}
