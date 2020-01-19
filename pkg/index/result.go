package index

// Result presents found game entries plus additional infromation, including ignored platforms requested etc.
type Result struct {
	Hits             []GameHit
	IgnoredPlatforms []string
}
