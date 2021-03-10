package json

// Project groups configuration properties of the whole GameDBV project.
type Project struct {
	API       API                 `json:"API"`
	Database  Database            `json:"Database"`
	Directory string              `json:"Directory"`
	Platforms map[string]Platform `json:"Platforms"`
}
