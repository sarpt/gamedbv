package parser

// Config provides information about files to be parsed by xml
type Config interface {
	Filepath() string
}

// ModelProvider returns pointer to data structure that should be filled by parser
type ModelProvider interface {
	Model() interface{}
}
