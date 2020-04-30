package index

// Searcher is used for searching in the indexes of the same type, eg. all bleve indexes could be aggregated to return batch of results
type Searcher interface {
	Search(SearchParameters) (Result, error)
	AddIndex(Config) error
	Close() error
}

// Creator is responsible for new index creation
type Creator interface {
	CreateIndex(string, []GameSource) error
}
