package bleve

import (
	"math"

	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/mapping"

	"github.com/sarpt/gamedbv/pkg/index"
)

const defaultBatchLength = 1000

// Creator satisifies interface of the same name in index package. Hanldes creation of bleve index
type Creator struct{}

// CreateIndex saves parsed index on the disk. The function uses batching to speed up the indexing
func (c Creator) CreateIndex(docType string, filepath string, games []index.GameSource) error {
	mapping := createIndexMapping(docType)

	index, err := bleve.New(filepath, mapping)
	if err != nil {
		return err
	}

	numberOfBatches := getNumberOfBatches(defaultBatchLength, len(games))
	for i := 1; i <= numberOfBatches; i++ {
		firstIdxToBatch := (i - 1) * defaultBatchLength
		lastIdxToBatch := i * defaultBatchLength
		if lastIdxToBatch > len(games) {
			lastIdxToBatch = len(games)
		}

		batch := index.NewBatch()

		gamesToBatch := games[firstIdxToBatch:lastIdxToBatch]
		for _, game := range gamesToBatch {
			idxSource := Data{
				Name:    game.GetName(),
				Region:  game.GetRegion(),
				docType: docType,
			}
			err = batch.Index(game.GetID(), idxSource)
			if err != nil {
				return err
			}
		}

		err = index.Batch(batch)
		if err != nil {
			return err
		}
	}

	return nil
}

func createIndexMapping(docType string) *mapping.IndexMappingImpl {
	mapping := bleve.NewIndexMapping()

	gametdbDocMapping := createGameTDBDocumentMapping()
	mapping.AddDocumentMapping(docType, gametdbDocMapping)

	return mapping
}

func createGameTDBDocumentMapping() *mapping.DocumentMapping {
	docMapping := bleve.NewDocumentMapping()

	nameField := bleve.NewTextFieldMapping()
	nameField.Store = true
	nameField.IncludeInAll = true
	nameField.Index = true
	docMapping.AddFieldMappingsAt("Name", nameField)

	textField := bleve.NewTextFieldMapping()
	textField.Store = false
	textField.IncludeInAll = false
	textField.Index = true
	docMapping.AddFieldMappingsAt("Region", textField)

	return docMapping
}

func getNumberOfBatches(batchLength int, collectionLength int) int {
	return int(math.Ceil(float64(collectionLength) / float64(batchLength)))
}
