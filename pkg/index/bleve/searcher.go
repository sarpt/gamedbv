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
func (s *Searcher) AddIndex(conf index.Config) error {
	indexPath, err := conf.GetIndexFilePath()
	if err != nil {
		return err
	}

	_, err = os.Stat(indexPath)
	if err != nil {
		return err
	}

	index, err := bleve.Open(indexPath)
	if err != nil {
		return err
	}

	s.indexes[conf.GetPlatform()] = index
	return nil
}

// Search returns aggregated results from indexes added to alias
func (s Searcher) Search(params index.SearchParameters) (string, error) {
	indexAlias := bleve.NewIndexAlias()
	for _, plat := range params.Platforms {
		if idx, ok := s.indexes[plat]; ok {
			indexAlias.Add(idx)
		} else {
			return "", fmt.Errorf("Could not search for platform: " + plat)
		}
	}

	query := bleve.NewConjunctionQuery()

	textQuery := bleve.NewMatchQuery(params.Text)
	query.AddQuery(textQuery)

	if len(params.Regions) > 0 {
		anyRegionQuery := bleve.NewDisjunctionQuery()

		for _, region := range params.Regions {
			regionQuery := bleve.NewTermQuery(region)
			anyRegionQuery.AddQuery(regionQuery)
		}

		query.AddQuery(anyRegionQuery)
	}

	request := bleve.NewSearchRequest(query)
	request.Fields = []string{nameField}
	request.Size = maxNumberOfResults

	result, err := indexAlias.Search(request)
	if err != nil {
		return "", err
	}

	var hits string
	for _, hit := range result.Hits {
		for key, value := range hit.Fields {
			if key == nameField {
				hits = fmt.Sprintln(hits + value.(string))
			}
		}
	}

	return hits, nil
}

// NewSearcher initializes bleve implementation of Searcher
func NewSearcher(configs []index.Config) (*Searcher, []index.Config) {
	var ignored []index.Config
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