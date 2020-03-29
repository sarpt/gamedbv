package index

// Description provides information about
type Description struct {
	Synopsis string
}

// GameSource provides information neccessary for game indexing
type GameSource struct {
	UID          string
	Name         string
	Descriptions []Description
}
