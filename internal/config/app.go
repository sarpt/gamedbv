package config

import (
	"os"
	"path"

	"github.com/sarpt/gamedbv/internal/api"
	"github.com/sarpt/gamedbv/internal/config/json"
	"github.com/sarpt/gamedbv/internal/dl"
	"github.com/sarpt/gamedbv/internal/games"
	"github.com/sarpt/gamedbv/internal/idx"
	"github.com/sarpt/gamedbv/pkg/db"
	"github.com/sarpt/gamedbv/pkg/index"
	"github.com/sarpt/gamedbv/pkg/platform"
)

// App groups configuration properties of the whole GameDBV project
// The app is used as an adapter for configurations required by all of the GameDBV packages and components
type App struct {
	directoryPath string
	platforms     map[string]Platform
	database      db.Config
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
		database: db.Config{
			Variant:  json.DefaultConfig.Database.Variant,
			Path:     path.Join(directoryPath, json.DefaultConfig.Database.FileName),
			MaxLimit: json.DefaultConfig.Database.MaxLimit,
		},
	}

	newApp.platforms = make(map[string]Platform)
	for variant, plat := range json.DefaultConfigsPerPlatform {
		platformDirectory := path.Join(directoryPath, plat.Directory)
		newApp.platforms[variant] = Platform{
			name:    plat.Name,
			dirPath: platformDirectory,
			source: Source{
				archiveFilepath: path.Join(platformDirectory, plat.Source.ArchiveFilename),
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

// platform returns platform config
func (cfg App) platform(variant platform.Variant) Platform {
	return cfg.platforms[variant.ID()]
}

// Database returns information neccessary to connect to persistence
func (cfg App) Database() db.Config {
	return cfg.database
}

// Idx returns confiuration for Idx component
func (cfg App) Idx(variant platform.Variant) idx.Config {
	platformConfig := cfg.platform(variant)

	return idx.Config{
		IndexFilepath:   platformConfig.index.path,
		IndexVariant:    platformConfig.index.variant,
		Name:            platformConfig.name,
		DocType:         platformConfig.index.docType,
		SourceFilename:  platformConfig.source.filename,
		SourceFilepath:  platformConfig.source.filepath,
		ArchiveFilepath: platformConfig.source.archiveFilepath,
	}
}

// Games returns configuration for Games component
func (cfg App) Games() games.Config {
	indexes := make(map[platform.Variant]index.Config)

	for platID, plat := range cfg.platforms {
		variant, err := platform.Get(platID)
		if err != nil {
			continue
		}

		indexes[variant] = index.Config{
			Filepath: plat.index.path,
			Variant:  plat.index.variant,
			Name:     plat.name,
			DocType:  plat.index.docType,
		}

	}

	return games.Config{
		Database: cfg.database,
		Indexes:  indexes,
	}
}

// API returns configuration for Api component
func (cfg App) API() api.Config {
	return api.Config{
		GamesConfig: cfg.Games(),
	}
}

// Dl returns configuration for Dl component
func (cfg App) Dl(variant platform.Variant) dl.Config {
	platformConfig := cfg.platform(variant)

	return dl.Config{
		DirectoryPath:   platformConfig.dirPath,
		Filepath:        platformConfig.source.archiveFilepath,
		ForceRedownload: platformConfig.source.forceDownload,
		URL:             platformConfig.source.url,
		PlatformName:    platformConfig.name,
	}
}
