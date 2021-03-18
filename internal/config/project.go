package config

import (
	"os"
	"path"
	"time"

	"github.com/sarpt/gamedbv/internal/api"
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
	Database  db.Config
	Directory string
	API       api.Config
	Dl        dl.Config
	Games     games.Config
	Idx       idx.Config
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

	var jsonProject jsonConfig.Project
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
	newApp.Database = db.Config{
		MaxLimit: jsonProject.Database.MaxLimit,
		Path:     path.Join(jsonProject.Directory, jsonProject.Database.Filename),
		Variant:  jsonProject.Database.Variant,
	}

	newApp.createDl(jsonProject)
	newApp.createIdx(jsonProject)
	newApp.createGamesConfig(jsonProject)

	err = newApp.createAPIConfig(jsonProject)

	return *newApp, err
}

func (cfg *Project) createDl(jsonApp jsonConfig.Project) {
	sources := map[platform.Variant]dl.SourceConfig{}

	for platformName := range jsonApp.Platforms {
		variant, err := platform.Get(platformName)
		if err != nil {
			continue // TODO: report from createDl function and handle
		}

		platformConfig := cfg.platform(variant)
		platformDirectoryPath := path.Join(cfg.Directory, platformConfig.Directory)

		sources[variant] = dl.SourceConfig{
			DirectoryPath:   platformDirectoryPath,
			Filepath:        path.Join(platformDirectoryPath, platformConfig.Source.ArchiveFilename),
			ForceRedownload: platformConfig.Source.ForceDownload,
			URL:             platformConfig.Source.URL,
			PlatformName:    platformConfig.Name,
		}
	}

	cfg.Dl = dl.Config{
		Sources: sources,
		Address: jsonApp.API.DlRPCAddress,
		Port:    jsonApp.API.DlRPCPort,
	}
}

func (cfg *Project) createIdx(jsonApp jsonConfig.Project) {
	indexes := map[platform.Variant]idx.IndexConfig{}

	for platformName := range jsonApp.Platforms {
		variant, err := platform.Get(platformName)
		if err != nil {
			continue // TODO: report from createIdx function and handle
		}

		platformConfig := cfg.platform(variant)
		platformDirectoryPath := path.Join(cfg.Directory, platformConfig.Directory)

		indexes[variant] = idx.IndexConfig{
			IndexFilepath:   path.Join(platformDirectoryPath, platformConfig.Index.Directory),
			IndexVariant:    platformConfig.Index.Variant,
			Name:            platformConfig.Name,
			DocType:         platformConfig.Index.DocType,
			SourceFilename:  platformConfig.Source.Filename,
			SourceFilepath:  path.Join(platformDirectoryPath, platformConfig.Source.Filename),
			ArchiveFilepath: path.Join(platformDirectoryPath, platformConfig.Source.ArchiveFilename),
		}
	}

	cfg.Idx = idx.Config{
		Address: jsonApp.API.IdxRPCAddress,
		Port:    jsonApp.API.IdxRPCAddress,
		Indexes: indexes,
	}
}

// platform returns platform config.
func (cfg Project) platform(variant platform.Variant) jsonConfig.Platform {
	return cfg.platforms[variant.ID()]
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
