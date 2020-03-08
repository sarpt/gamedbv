package bleve

import (
	"fmt"
	"os"

	"github.com/blevesearch/bleve"
	"github.com/sarpt/gamedbv/pkg/index"
)

const maxNumberOfResults = 1000
const nameField = "Name"

// Searcher implements the interface of the same name for indexes created by bleve
type Searcher struct {
	indexes map[string]bleve.Index
}

// AddIndex uses platform config to add another index to be used during searching
func (s *Searcher) AddIndex(conf index.PlatformConfig) error {
	indexPath := conf.IndexFilepath()
	_, err := os.Stat(indexPath)
	if err != nil {
		return err
	}

	index, err := bleve.Open(indexPath)
	if err != nil {
		return err
	}

	s.indexes[conf.Name()] = index
	return nil
}

// Search returns aggregated results from indexes added to alias
func (s Searcher) Search(params index.SearchParameters) (index.Result, error) {
	searchResult := index.Result{}

	indexAlias := bleve.NewIndexAlias()
	for _, plat := range params.Platforms {
		if idx, ok := s.indexes[plat]; ok {
			indexAlias.Add(idx)
		} else {
			searchResult.IgnoredPlatforms = append(searchResult.IgnoredPlatforms, plat)
		}
	}

	if len(searchResult.IgnoredPlatforms) == len(params.Platforms) {
		return searchResult, fmt.Errorf("Could not execute search due to lack of indexes for any of the provided platforms")
	}

	query := bleve.NewDisjunctionQuery()
	query.SetMin(1)

	textQuery := bleve.NewMatchPhraseQuery(params.Text)
	query.AddQuery(textQuery)

	prefixQuery := bleve.NewPrefixQuery(params.Text)
	query.AddQuery(prefixQuery)

	request := bleve.NewSearchRequest(query)
	request.Size = maxNumberOfResults

	result, err := indexAlias.Search(request)
	if err != nil {
		return searchResult, err
	}

	for _, hit := range result.Hits {
		gameHit := index.GameHit{
			ID: hit.ID,
		}
		searchResult.Hits = append(searchResult.Hits, gameHit)
	}

	return searchResult, nil
}

// NewSearcher initializes bleve implementation of Searcher
func NewSearcher(configs []index.PlatformConfig) (*Searcher, []index.PlatformConfig) {
	var ignored []index.PlatformConfig
	s := &Searcher{
		indexes: make(map[string]bleve.Index),
	}

	for _, conf := range configs {
		err := s.AddIndex(conf)
		if err != nil {
			ignored = append(ignored, conf)
		}
	}

	return s, ignored
}
