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
