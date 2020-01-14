package bleve

import (
	"math"

	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/mapping"

	"github.com/sarpt/gamedbv/pkg/gametdb"
)

const defaultBatchLength = 1000

// Creator satisifies interface of the same name in index package. Hanldes creation of bleve index
type Creator struct{}

// CreateIndex saves parsed index on the disk. The function uses batching to speed up the indexing
func (c Creator) CreateIndex(filepath string, games []gametdb.Game) error {
	mapping := createIndexMapping()

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
			err = batch.Index(game.ID, game)
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

func createIndexMapping() *mapping.IndexMappingImpl {
	mapping := bleve.NewIndexMapping()
	gametdbDocMapping := createGameTDBDocumentMapping()

	mapping.AddDocumentMapping(gametdb.GameDoctype, gametdbDocMapping)

	return mapping
}

func createGameTDBDocumentMapping() *mapping.DocumentMapping {
	docmapping := bleve.NewDocumentMapping()

	nameField := bleve.NewTextFieldMapping()
	nameField.Store = true
	nameField.IncludeInAll = false
	nameField.Index = true
	docmapping.AddFieldMappingsAt("Name", nameField)

	textField := bleve.NewTextFieldMapping()
	textField.Store = false
	textField.IncludeInAll = false
	textField.Index = true
	docmapping.AddFieldMappingsAt("Region", textField)

	return docmapping
}

func getNumberOfBatches(batchLength int, collectionLength int) int {
	return int(math.Ceil(float64(collectionLength) / float64(batchLength)))
}
