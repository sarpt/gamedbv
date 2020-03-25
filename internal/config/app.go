package config

import (
	"os"
	"path"

	"github.com/sarpt/gamedbv/internal/config/json"
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

// NewApp returns new config
// At the moment, it only copies the defaults - reading from environment/overrides to be provided
func NewApp() (App, error) {
	var newApp App

	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return newApp, err
	}

	directoryPath := path.Join(userConfigDir, json.DefaultConfig.Directory)
	newApp = App{
		directoryPath: directoryPath,
		database: Database{
			variant:  json.DefaultConfig.Database.Variant,
			path:     path.Join(directoryPath, json.DefaultConfig.Database.FileName),
			maxLimit: json.DefaultConfig.Database.MaxLimit,
		},
	}

	newApp.platforms = make(map[string]Platform)
	for variant, plat := range json.DefaultConfigsPerPlatform {
		platformDirectory := path.Join(directoryPath, plat.Directory)
		newApp.platforms[variant] = Platform{
			name:    plat.Name,
			dirPath: platformDirectory,
			source: Source{
				archiveFilename: path.Join(platformDirectory, plat.Source.ArchiveFilename),
				filepath:        path.Join(platformDirectory, plat.Source.Filename),
				filename:        plat.Source.Filename,
				name:            plat.Name,
				archived:        plat.Source.Archived,
				format:          plat.Source.Format,
				forceDownload:   plat.Source.ForceDownload,
				url:             plat.Source.URL,
			},
			index: Index{
				path:    path.Join(platformDirectory, plat.Index.Directory),
				variant: plat.Index.Variant,
				docType: plat.Index.DocType,
			},
		}
	}

	return newApp, nil
}
