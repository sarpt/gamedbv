package config

// Index includes information about variant of index (bleve, solr etc), stored directory etc.
type Index struct {
	variant string
	path    string
	docType string
}
