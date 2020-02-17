package db

// Config contains information neccessary for db access
type Config interface {
	DatabasePath() string
	DatabaseVariant() string
}
