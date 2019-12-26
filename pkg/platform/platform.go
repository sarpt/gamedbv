package platform

import "strings"

const (
	// Wii is Nintendo Wii platform
	Wii = "wii"
	// Wiiu is Nintendo WiiU platform
	Wiiu = "wiiu"
	// Ps3 is PlayStation 3 platform
	Ps3 = "ps3"
	// Nds is Nintendo DS platform
	Nds = "nds"
	// N3ds is Nintendo 3DS platform
	N3ds = "n3ds"
	// Switch is Nintendo Switch platfrom
	Switch = "switch"
)

// Variant is used to specify type of Database to download.
type Variant struct {
	value string
}

var platforms = []string{Wii, Wiiu, Ps3, Nds, N3ds, Switch}

func (variant *Variant) String() string {
	return variant.value
}

// Set validates and sets platform if it's a correct string value or throws an error when it's a not supported value
func (variant *Variant) Set(value string) error {
	lowerCaseValue := strings.ToLower(value)

	if !isValueCorrectPlatform(lowerCaseValue) {
		return &IncorrectPlatformError{incorrectPlatform: value}
	}

	variant.value = lowerCaseValue
	return nil
}

// IsSet returns true when the variant has been correctly set
func (variant Variant) IsSet() bool {
	return len(variant.value) > 0
}

// GetAllPlatforms returns all possible variants of platform databases that could be downloaded
func GetAllPlatforms() []Variant {
	var allPlatforms []Variant

	for _, platform := range platforms {
		allPlatforms = append(allPlatforms, Variant{value: platform})
	}

	return allPlatforms
}

func isValueCorrectPlatform(value string) bool {
	for _, platform := range platforms {
		if value == platform {
			return true
		}
	}

	return false
}
