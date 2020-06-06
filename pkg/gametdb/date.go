package gametdb

// Date presents xml date node in GameTDB file
// It's represents release date of a game
type Date struct {
	Year  uint `xml:"year,attr"`
	Month uint `xml:"month,attr"`
	Day   uint `xml:"day,attr"`
}
