package json

// Platform groups information used for platform database handling
type Platform struct {
	Directory string
	Name      string
	Source    Source
	Index     Index
}
