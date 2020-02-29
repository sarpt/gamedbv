package config

// Database represents information about persistence instance used for holding platform data
type Database struct {
	path    string
	variant string
}

// Path return path of the database which persists information about entries parsed from the source file
func (conf Database) Path() string {
	return conf.path
}

// Variant return path of the database which persists information about entries parsed from the source file
func (conf Database) Variant() string {
	return conf.variant
}
