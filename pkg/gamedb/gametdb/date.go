package gametdb

// Date presents xml date node in GameTDB file
// It's represents release date of a game
type Date struct {
	Year  int `xml:"year,attr"`
	Month int `xml:"month,attr"`
	Day   int `xml:"day,attr"`
}
