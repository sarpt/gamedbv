package config

import (
	"github.com/sarpt/gamedbv/pkg/platform"
)

// App groups configuration properties of the whole GameDBV project
type App struct {
	directoryPath string
	platforms     map[string]Platform
	database      Database
}

// GetBaseDirectoryPath returns absolute path of GameDBV directory
func (conf App) GetBaseDirectoryPath() string {
	return conf.directoryPath
}

// Platform returns platform config
func (conf App) Platform(variant platform.Variant) Platform {
	return conf.platforms[variant.String()]
}

// Database returns information neccessary to connect to persistence
func (conf App) Database() Database {
	return conf.database
}
