package index

// SearchParameters provide information what criteria for results are expected
type SearchParameters struct {
	Text      string
	Regions   []string
	Platforms []string
}
