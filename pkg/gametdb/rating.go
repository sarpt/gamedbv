package gametdb

// Rating presents xml rating node in GameTDB file
// It presents rating of the game in respective age rating system
type Rating struct {
	System string `xml:"type,attr"`
	Value  string `xml:"value,attr"`
}
