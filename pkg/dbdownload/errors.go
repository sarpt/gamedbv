package dbdownload

// IncorrectPlatformError is thrown when incorrect platform is set.
type IncorrectPlatformError struct {
	incorrectPlatform string
}

func (err *IncorrectPlatformError) Error() string {
	return "Incorrect platform: " + err.incorrectPlatform
}
