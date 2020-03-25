package config

// Database represents information about persistence instance used for holding platform data
type Database struct {
	path     string
	variant  string
	maxLimit int
}

// Path returns path of the database which persists information about entries parsed from the source file
func (conf Database) Path() string {
	return conf.path
}

// Variant returns path of the database which persists information about entries parsed from the source file
func (conf Database) Variant() string {
	return conf.variant
}

// MaxLimit returns the maximum number of elements that should returned per one SELECT
// This parameter is introduced due to SQLite default (hard?) limit of parameters in expression equal to 999
// Due to problem of GORM using IN keyword while preloading and selecting related tables, limit is used to work around this issue
func (conf Database) MaxLimit() int {
	return conf.maxLimit
}
