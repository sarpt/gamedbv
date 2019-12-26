package gametdb

// Input presents xml input node in GameTDB file
// It presents details about input features game might posses
type Input struct {
	Players  string    `xml:"players,attr"`
	Controls []Control `xml:"control"`
}
