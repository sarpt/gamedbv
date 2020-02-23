package db

// Config contains information neccessary for db access
type Config interface {
	Path() string
	Variant() string
}
