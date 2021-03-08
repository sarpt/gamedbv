package json

// Platform groups information used for platform database handling
type Platform struct {
	Directory string `json:"Directory"`
	Name      string `json:"Name"`
	Source    Source `json:"Source"`
	Index     Index  `json:"Index"`
}
