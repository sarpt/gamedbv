package json

// Index includes information about variant of index (bleve, solr etc), stored directory etc.
type Index struct {
	Variant   string `json:"Variant"`
	Directory string `json:"Directory"`
	DocType   string `json:"DocType"`
}
