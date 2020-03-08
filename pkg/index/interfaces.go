package index

// PlatformConfig provides index information and settings
type PlatformConfig interface {
	IndexFilepath() string
	IndexVariant() string
	Name() string
	DocType() string
}

// Searcher is used for searching in the indexes of the same type, eg. all bleve indexes could be aggregated to return batch of results
type Searcher interface {
	Search(SearchParameters) (Result, error)
	AddIndex(PlatformConfig) error
}

// Creator is responsible for new index creation
type Creator interface {
	CreateIndex(string, []GameSource) error
}
