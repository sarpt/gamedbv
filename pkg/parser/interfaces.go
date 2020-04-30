package parser

// ModelProvider returns pointer to data structure that should be filled by parser
type ModelProvider interface {
	Model() interface{}
}
