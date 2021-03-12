package config

import (
	"os"
	"path"
	"time"

	"github.com/sarpt/gamedbv/internal/api"
	"github.com/sarpt/gamedbv/internal/config/json"
	jsonConfig "github.com/sarpt/gamedbv/internal/config/json"
	"github.com/sarpt/gamedbv/internal/dl"
	"github.com/sarpt/gamedbv/internal/games"
	"github.com/sarpt/gamedbv/internal/idx"
	"github.com/sarpt/gamedbv/pkg/db"
	"github.com/sarpt/gamedbv/pkg/index"
	"github.com/sarpt/gamedbv/pkg/platform"
)

// Project groups configuration properties of the whole GameDBV project.
type Project struct {
	Directory string
	API       api.Config
	Games     games.Config
	Database  db.Config
	platforms map[string]jsonConfig.Platform
}

// Create returns GameDBV project config.
// In-case the override for config does not exist in user's config directory, the function saves the bundled during compile-time default.
// Otherwise the function tries to read user provided one, without any manipulation to it.
func Create() (Project, error) {
	newApp := &Project{}

	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return *newApp, err
	}

	var jsonProject json.Project
	if userConfigExists() {
		jsonProject, err = readUserConfig()
		if err != nil {
			return *newApp, err
		}
	} else {
		jsonProject, err = jsonConfig.Default()
		if err != nil {
			return *newApp, err
		}

		err = writeUserConfig(jsonProject)
		if err != nil {
			return *newApp, err
		}
	}

	newApp.platforms = jsonProject.Platforms
	newApp.Directory = path.Join(userHomeDir, jsonProject.Directory)

	newApp.createDatabaseConfig(jsonProject)
	newApp.createGamesConfig(jsonProject)

	err = newApp.createAPIConfig(jsonProject)

	return *newApp, err
}

// Dl returns configuration for Dl component.
func (cfg Project) Dl(variant platform.Variant) dl.Config {
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
func (cfg Project) Idx(variant platform.Variant) idx.Config {
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
func (cfg Project) platform(variant platform.Variant) jsonConfig.Platform {
	return cfg.platforms[variant.ID()]
}

func (cfg *Project) createDatabaseConfig(jsonApp jsonConfig.Project) {
	cfg.Database = db.Config{
		MaxLimit: jsonApp.Database.MaxLimit,
		Path:     path.Join(cfg.Directory, jsonApp.Database.Filename),
		Variant:  jsonApp.Database.Variant,
	}
}

func (cfg *Project) createGamesConfig(jsonApp jsonConfig.Project) {
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

func (cfg *Project) createAPIConfig(jsonApp jsonConfig.Project) error {
	readTimeout, err := time.ParseDuration(jsonApp.API.ReadTimeout)
	if err != nil {
		return err
	}

	writeTimeout, err := time.ParseDuration(jsonApp.API.WriteTimeout)
	if err != nil {
		return err
	}

	rpcDialTimeout, err := time.ParseDuration(jsonApp.API.RPCDialTimeout)
	if err != nil {
		return err
	}

	cfg.API = api.Config{
		Debug:          jsonApp.API.Debug,
		DlRPCAddress:   jsonApp.API.DlRPCAddress,
		DlRPCPort:      jsonApp.API.DlRPCPort,
		GamesConfig:    cfg.Games,
		IPAddress:      jsonApp.API.IPAddress,
		NetInterface:   jsonApp.API.NetInterface,
		Port:           jsonApp.API.Port,
		ReadTimeout:    readTimeout,
		RPCDialTimeout: rpcDialTimeout,
		WriteTimeout:   writeTimeout,
	}

	return nil
}
