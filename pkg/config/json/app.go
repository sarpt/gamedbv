package json

// App groups configuration properties of the whole GameDBV project
type App struct {
	Directory string
	Platforms map[string]Platform
	Database  Database
}
