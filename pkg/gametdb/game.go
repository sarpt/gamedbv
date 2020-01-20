package gametdb

// Game presents xml game node in GameTDB file
// It's a main database entry about a specific game
type Game struct {
	Name         string   `xml:"name,attr"`
	ID           string   `xml:"id"`
	Region       string   `xml:"region"`
	Languages    string   `xml:"languages"`
	Locales      []Locale `xml:"locale"`
	Developer    string   `xml:"developer"`
	Publisher    string   `xml:"publisher"`
	Date         Date     `xml:"date"`
	Genre        string   `xml:"genre"`
	Rating       Rating   `xml:"rating"`
	WiFi         WiFi     `xml:"wi-fi"`
	Input        Input    `xml:"input"`
	Rom          Rom      `xml:"rom"`
	PlatformType string   `xml:"type"`
}

// GetID returns ID of the game - usually a unique serial number that differs in format depending of platform
func (g Game) GetID() string {
	return g.ID
}

// GetName returns a game name
func (g Game) GetName() string {
	return g.Name
}

// GetRegion returns a region in which game was released
func (g Game) GetRegion() string {
	return g.Region
}
