package gametdb

// Datafile is a root xml element of GameTDB
type Datafile struct {
	Games []Game `xml:"game"`
}
