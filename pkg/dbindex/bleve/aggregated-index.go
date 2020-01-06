package bleve

import (
	"fmt"
	"os"

	"github.com/blevesearch/bleve"
	"github.com/sarpt/gamedbv/pkg/dbindex/shared"
	"github.com/sarpt/gamedbv/pkg/platform"
)

// AggregatedIndex implements the interface of the same name for indexes created by bleve
type AggregatedIndex struct {
	indexAlias bleve.IndexAlias
}

// Add uses platform config to add another index to be used during searching
func (aidx AggregatedIndex) Add(conf platform.Config) error {
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

	aidx.indexAlias.Add(index)
	return nil
}

// Search returns aggregated results from indexes added to alias
func (aidx AggregatedIndex) Search(params shared.SearchParameters) (string, error) {
	query := bleve.NewMatchQuery(params.Text)
	request := bleve.NewSearchRequest(query)
	request.Fields = []string{"Name"}
	request.Size = 1000

	result, err := aidx.indexAlias.Search(request)
	if err != nil {
		return "", err
	}

	var hits string
	for _, hit := range result.Hits {
		fields := hit.Fields
		for key, value := range fields {
			if key == "Name" {
				hits = fmt.Sprintln(hits + value.(string))
			}
		}
	}

	return hits, nil
}

// New initializes bleve implementation of AggregatedIndex
func New() AggregatedIndex {
	return AggregatedIndex{
		indexAlias: bleve.NewIndexAlias(),
	}
}
