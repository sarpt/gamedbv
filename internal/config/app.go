package config

import (
	"os"
	"path"
	"time"

	"github.com/sarpt/gamedbv/internal/api"
	"github.com/sarpt/gamedbv/internal/config/json"
	"github.com/sarpt/gamedbv/internal/dl"
	"github.com/sarpt/gamedbv/internal/games"
	"github.com/sarpt/gamedbv/internal/idx"
	"github.com/sarpt/gamedbv/pkg/db"
	"github.com/sarpt/gamedbv/pkg/index"
	"github.com/sarpt/gamedbv/pkg/platform"
)

// App groups configuration properties of the whole GameDBV project.
type App struct {
	Directory string
	API       api.Config
	Games     games.Config
	Database  db.Config
	platforms map[string]json.Platform
}

// NewApp returns new config.
// At the moment it only reads from the default embedded at compile-time.
func NewApp() (App, error) {
	newApp := &App{}

	userDir, err := os.UserHomeDir()
	if err != nil {
		return *newApp, err
	}

	defaultApp, err := json.DefaultApp()
	if err != nil {
		return *newApp, err
	}

	newApp.platforms = defaultApp.Platforms
	newApp.Directory = path.Join(userDir, defaultApp.Directory)

	newApp.createDatabaseConfig(defaultApp)
	newApp.createGamesConfig(defaultApp)

	err = newApp.createAPIConfig(defaultApp)

	return *newApp, err
}

// Dl returns configuration for Dl component.
func (cfg App) Dl(variant platform.Variant) dl.Config {
	platformConfig := cfg.platform(variant)
	platformDirectoryPath := path.Join(cfg.Directory, platformConfig.Directory)

	return dl.Config{
		DirectoryPath:   platformDirectoryPath,
		Filepath:        path.Join(platformDirectoryPath, platformConfig.Source.ArchiveFilename),
		ForceRedownload: platformConfig.Source.ForceDownload,
		URL:             platformConfig.Source.URL,
		PlatformName:    platformConfig.Name,
	}
}

// Idx returns confiuration for Idx component.
func (cfg App) Idx(variant platform.Variant) idx.Config {
	platformConfig := cfg.platform(variant)
	platformDirectoryPath := path.Join(cfg.Directory, platformConfig.Directory)

	return idx.Config{
		IndexFilepath:   path.Join(platformDirectoryPath, platformConfig.Index.Directory),
		IndexVariant:    platformConfig.Index.Variant,
		Name:            platformConfig.Name,
		DocType:         platformConfig.Index.DocType,
		SourceFilename:  platformConfig.Source.Filename,
		SourceFilepath:  path.Join(platformDirectoryPath, platformConfig.Source.Filename),
		ArchiveFilepath: path.Join(platformDirectoryPath, platformConfig.Source.ArchiveFilename),
	}
}

// platform returns platform config.
func (cfg App) platform(variant platform.Variant) json.Platform {
	return cfg.platforms[variant.ID()]
}

func (cfg *App) createDatabaseConfig(jsonApp json.App) {
	cfg.Database = db.Config{
		MaxLimit: jsonApp.Database.MaxLimit,
		Path:     path.Join(cfg.Directory, jsonApp.Database.Filename),
		Variant:  jsonApp.Database.Variant,
	}
}

func (cfg *App) createGamesConfig(jsonApp json.App) {
	indexes := make(map[platform.Variant]index.Config)

	for platID, platformConfig := range jsonApp.Platforms {
		variant, err := platform.Get(platID)
		if err != nil {
			continue
		}

		platformDirectoryPath := path.Join(cfg.Directory, platformConfig.Directory)
		indexes[variant] = index.Config{
			Filepath: path.Join(platformDirectoryPath, platformConfig.Index.Directory),
			Variant:  platformConfig.Index.Variant,
			Name:     platformConfig.Name,
			DocType:  platformConfig.Index.DocType,
		}

	}

	cfg.Games = games.Config{
		Database: cfg.Database,
		Indexes:  indexes,
	}
}

func (cfg *App) createAPIConfig(jsonApp json.App) error {
	apiReadTimeout, err := time.ParseDuration(jsonApp.API.ReadTimeout)
	if err != nil {
		return err
	}

	apiWriteTimeout, err := time.ParseDuration(jsonApp.API.WriteTimeout)
	if err != nil {
		return err
	}

	cfg.API = api.Config{
		Debug:        jsonApp.API.Debug,
		GamesConfig:  cfg.Games,
		IPAddress:    jsonApp.API.IPAddress,
		NetInterface: jsonApp.API.NetInterface,
		Port:         jsonApp.API.Port,
		ReadTimeout:  apiReadTimeout,
		WriteTimeout: apiWriteTimeout,
	}

	return nil
}
