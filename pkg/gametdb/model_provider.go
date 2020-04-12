package gametdb

// ModelProvider implements interface of the same name used for parsing gametdb source files
type ModelProvider struct {
	dataFile Datafile
}

// Model returns DataFile structure that can hold results of unmarshalling of gametdb xml files
func (mp *ModelProvider) Model() interface{} {
	return &mp.dataFile
}

// Root returns root xml element (Datafile)
func (mp *ModelProvider) Root() Datafile {
	return mp.dataFile
}

// Games returns games found in the source file
func (mp ModelProvider) Games() []Game {
	return mp.dataFile.Games
}
