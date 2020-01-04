package gametdb

// Locale presents xml local node in GameTDB db file
// Locale is responsible to present title and synposis of the game in specified Language
type Locale struct {
	Language string `xml:"lang,attr"`
	Title    string `xml:"title"`
	Synopsis string `xml:"synposis"`
}
