package json

// App groups configuration properties of the whole GameDBV project
type App struct {
	API       API
	Database  Database
	Directory string
	Platforms map[string]Platform
}
