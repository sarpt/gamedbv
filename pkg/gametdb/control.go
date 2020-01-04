package gametdb

// Control presents xml control node in GameTDB file
// It includes details about control inputs types
type Control struct {
	Controller string `xml:"type,attr"`
	Required   string `xml:"required,attr"`
}
