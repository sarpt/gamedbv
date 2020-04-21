package platform

import "strings"

// Variant is used to specify type of Database to download.
type Variant struct {
	id   string
	name string
}

// Get returns Variant when supported platform variant exists with the id
// Otherwise IncorrectPlatformError is returned
func Get(id string) (Variant, error) {
	variant, ok := variants[strings.ToLower(id)]

	if !ok {
		return variant, &IncorrectPlatformError{incorrectPlatform: id}
	}

	return variant, nil
}

// ID returns platform variant id
func (variant Variant) ID() string {
	return variant.id
}

// Name returns platform variant name as commercially known
func (variant Variant) Name() string {
	return variant.name
}

func (variant Variant) String() string {
	return variant.Name()
}
