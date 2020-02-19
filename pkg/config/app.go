package config

import (
	"os"
	"path"
)

// App groups configuration properties of the whole GameDBV project
type App struct {
	BaseDirectory string
}

// GetBaseDirectoryPath returns absolute path of GameDBV directory
func (conf App) GetBaseDirectoryPath() (string, error) {
	homePath, err := os.UserConfigDir()

	return path.Join(homePath, conf.BaseDirectory), err
}
