package bleve

import (
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/mapping"
)

func createIndexMapping() *mapping.IndexMappingImpl {
	mapping := bleve.NewIndexMapping()

	gametdbDocMapping := createGameTDBDocumentMapping()
	mapping.DefaultMapping = gametdbDocMapping

	return mapping
}

func createGameTDBDocumentMapping() *mapping.DocumentMapping {
	docMapping := bleve.NewDocumentMapping()

	nameField := bleve.NewTextFieldMapping()
	nameField.Store = true
	nameField.IncludeInAll = true
	nameField.Index = true
	docMapping.AddFieldMappingsAt("Name", nameField)

	descriptionDocMapping := createDescriptionDocumentMapping()
	docMapping.AddSubDocumentMapping("Descriptions", descriptionDocMapping)

	return docMapping
}

func createDescriptionDocumentMapping() *mapping.DocumentMapping {
	descriptionMapping := bleve.NewDocumentMapping()

	textField := bleve.NewTextFieldMapping()
	textField.Store = false
	textField.IncludeInAll = true
	textField.Index = true
	descriptionMapping.AddFieldMappingsAt("Synopsis", textField)

	return descriptionMapping
}
