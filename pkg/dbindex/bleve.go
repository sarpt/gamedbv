package dbindex

import (
	"github.com/blevesearch/bleve/mapping"
	"github.com/sarpt/gamedbv/pkg/gametdb"

	"github.com/blevesearch/bleve"
)

func createIndex(variantName string, filepath string, games []gametdb.Game) error {
	mapping := createIndexMapping()

	index, err := bleve.New(filepath, mapping)
	if err != nil {
		return err
	}

	err = index.Index(variantName, games)
	return err
}

func createIndexMapping() *mapping.IndexMappingImpl {
	mapping := bleve.NewIndexMapping()
	gametdbDocMapping := createGameTDBDocumentMapping()

	mapping.AddDocumentMapping(gametdb.GameDoctype, gametdbDocMapping)

	return mapping
}

func createGameTDBDocumentMapping() *mapping.DocumentMapping {
	docmapping := bleve.NewDocumentMapping()

	textField := bleve.NewTextFieldMapping()
	docmapping.AddFieldMappingsAt("Name", textField)
	docmapping.AddFieldMappingsAt("ID", textField)
	docmapping.AddFieldMappingsAt("Region", textField)
	docmapping.AddFieldMappingsAt("Languages", textField)
	docmapping.AddFieldMappingsAt("Developer", textField)
	docmapping.AddFieldMappingsAt("Publisher", textField)
	docmapping.AddFieldMappingsAt("Genre", textField)
	docmapping.AddFieldMappingsAt("PlatformType", textField)

	return docmapping
}
