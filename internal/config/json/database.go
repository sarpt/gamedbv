package json

// Database represents information about persistence instance used for holding platform data
type Database struct {
	FileName string
	Variant  string
	MaxLimit int
}
