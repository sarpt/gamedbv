package bleve

import (
	"math"

	"github.com/blevesearch/bleve"

	"github.com/sarpt/gamedbv/pkg/index"
)

const defaultBatchLength = 1000

// BleveCreator satisifies interface of the same name in index package. Hanldes creation of bleve index
type BleveCreator struct{}

// CreateIndex saves parsed index on the disk. The function uses batching to speed up the indexing
func (c BleveCreator) CreateIndex(filepath string, games []index.GameSource) error {
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
			err = batch.Index(game.UID, game)
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

func getNumberOfBatches(batchLength int, collectionLength int) int {
	return int(math.Ceil(float64(collectionLength) / float64(batchLength)))
}
