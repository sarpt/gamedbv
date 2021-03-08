package json

// Database represents information about persistence instance used for holding platform data
type Database struct {
	Filename string `json:"Filename"`
	Variant  string `json:"Variant"`
	MaxLimit int    `json:"MaxLimit"`
}
