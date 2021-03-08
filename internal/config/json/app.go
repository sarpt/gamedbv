package json

// App groups configuration properties of the whole GameDBV project
type App struct {
	API       API                 `json:"API"`
	Database  Database            `json:"Database"`
	Directory string              `json:"Directory"`
	Platforms map[string]Platform `json:"Platforms"`
}
