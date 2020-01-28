package db

// Checksum contains information neccessary to properly check the integrity of a file
type Checksum interface {
	Variant() string
	Value() string
}

// Rom provides information about known dumps of the game
type Rom interface {
	Version() string
	Name() string
	Size() string
	Checksums() []Checksum
}

// Language contains information about the code and the name of it
type Language interface {
	Code() string
	Name() string
}

// GameDescription contains detailed informations about games
type GameDescription interface {
	Language() Language
	Title() string
	Synopsis() string
}

// Game contains information about a single game release
type Game interface {
	ID() string
	Region() string
	Languages() []string
	Descriptions() []GameDescription
	Developer() string
	Publisher() string
	Roms() []Rom
	Date() string
	Rating() Rating
}

// Rating contains information about game rating (age restrictions, etc)
type Rating interface {
	Language() Language
	Variant() string
	Value() string
}

// Company includes information about business entities involved with game
type Company interface {
	Name() string
}

// Genre includes infromation that helps categorize type of a game
type Genre interface {
	Language() Language
	Title() string
	Subgenres() []Genre
}

// Source includes information about website/database which contributes information about platform
// Different sources may contribute variable amount of information
// It's expected that at the very least basic information about games will be provided
type Source interface {
	Variant() string
	URL() string
}

// PlatformProvider is used to update information in the database about a platform
// One provider includes information about a single platform, from a single source
// It provides information about games, companies, genres etc.
// Providers should be implemented by the source handlers, since it's expected that db will handle multiple possible sources for platforms
// Note: at the moment PlatformProvider is loosely modelled after GameTDB, however it may expand/change when support for more sources is implemented
type PlatformProvider interface {
	Platform() string
	Games() []Game
	Companies() []Company
	Ratings() []Rating
	Genres() []Genre
	Source() Source
}

// Config contains information neccessary for db access
type Config interface {
	Path() string
}
