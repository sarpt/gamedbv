package gametdb

// WiFi presents xml wi-fi node in GameTDB file
// It presents details about wifi connectivity features game might posses
type WiFi struct {
	Players string `xml:"players,attr"`
}
