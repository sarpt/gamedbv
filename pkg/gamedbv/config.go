package gamedbv

import (
	"os"
	"path"
)

// Config groups configuration properties of the whole GameDBV project
type Config struct {
	BaseDirectory string
}

// GetBaseDirectoryPath returns absolute path of GameDBV directory
func (conf Config) GetBaseDirectoryPath() (string, error) {
	homePath, err := os.UserConfigDir()

	return path.Join(homePath, conf.BaseDirectory), err
}
