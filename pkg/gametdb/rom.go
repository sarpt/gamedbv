package gametdb

// Rom presents xml rom node in GameTDB file
// It includes information about rom dumps related to the game
type Rom struct {
	Version string `xml:"version,attr"`
	Name    string `xml:"name,attr"`
	Size    int    `xml:"size,attr"`
	Crc     string `xml:"crc,attr"`
	Md5     string `xml:"md5,attr"`
	Sha1    string `xml:"sha1,attr"`
}
