package config

// Platform groups information used for platform database handling
type Platform struct {
	dirPath string
	name    string
	index   Index
	source  Source
}
