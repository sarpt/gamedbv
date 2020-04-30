package bleve

import (
	"fmt"
	"os"

	"github.com/blevesearch/bleve"
	"github.com/sarpt/gamedbv/pkg/index"
)

const maxNumberOfResults = 1000
const nameField = "Name"

// BleveSearcher implements the interface of the same name for indexes created by bleve
type BleveSearcher struct {
	indexes map[string]bleve.Index
}

// NewBleveSearcher initializes bleve implementation of Searcher
func NewBleveSearcher(configs []index.Config) (*BleveSearcher, []index.Config) {
	var ignored []index.Config
	s := &BleveSearcher{
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

// Close closes all indexes of a Searcher opened by AddIndex
// Always needs to be called as NewSearcher invokes AddIndex too
func (s *BleveSearcher) Close() error {
	var closeError error

	for _, idx := range s.indexes {
		err := idx.Close()
		if err != nil {
			closeError = err
		}
	}

	return closeError
}

// AddIndex uses platform config to add another index to be used during searching
func (s *BleveSearcher) AddIndex(conf index.Config) error {
	indexPath := conf.Filepath
	_, err := os.Stat(indexPath)
	if err != nil {
		return err
	}

	index, err := bleve.Open(indexPath)
	if err != nil {
		return err
	}

	s.indexes[conf.Name] = index
	return nil
}

// Search returns aggregated results from indexes added to alias
func (s BleveSearcher) Search(params index.SearchParameters) (index.Result, error) {
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
