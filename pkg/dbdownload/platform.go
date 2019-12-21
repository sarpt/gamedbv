package dbdownload

// Platform is used to specify type of Database to download.
type Platform struct {
	value string
}

var platforms = []string{"wii", "wiiu", "ps3", "nds", "3ds", "switch"}

func (platform *Platform) String() string {
	return platform.value
}

// Set validates and sets platform if it's a correct string value or throws an error when it's a not supported value
func (platform *Platform) Set(value string) error {
	if !isValueCorrectPlatform(value) {
		return &IncorrectPlatformError{incorrectPlatform: value}
	}

	platform.value = value
	return nil
}

// IsSet returns true when the variant has been correctly set
func (platform Platform) IsSet() bool {
	return len(platform.value) > 0
}

// GetAllPlatforms returns all possible variants of GameTDB platform databases that could be downloaded
func GetAllPlatforms() []Platform {
	allPlatforms := make([]Platform, 0)

	for _, platform := range platforms {
		allPlatforms = append(allPlatforms, Platform{value: platform})
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
