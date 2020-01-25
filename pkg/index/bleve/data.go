package bleve

// Data provides fields to be used for indexing
type Data struct {
	Name    string
	Region  string
	docType string
}

// Type returns doctype to be used for bleve document mapping
func (s Data) Type() string {
	return s.docType
}
