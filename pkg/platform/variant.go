package platform

import "strings"

// Variant holds information about a console platform variant
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

func ByNames(names []string) ([]Variant, error) {
	var platforms []Variant

	if len(names) == 0 {
		platforms = append(platforms, All()...)
	} else {
		for _, val := range names {
			variant, err := Get(val)
			if err != nil {
				return platforms, err
			}

			platforms = append(platforms, variant)
		}
	}

	return platforms, nil
}

// ID returns platform variant id, should be unique per platform
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
