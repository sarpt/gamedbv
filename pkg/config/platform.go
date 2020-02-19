package config

import (
	"path"

	"github.com/sarpt/gamedbv/pkg/platform"
)

// Platform groups information used for platform database handling
type Platform struct {
	appConfig App
	directory string
	name      string
	source    Source
	index     Index
	database  Database
}

// GetConfig returns final platform database information object taking into accounts default values and passed overrides of settings.
// Todo: implement passing overrides
func GetConfig(dbPlatform platform.Variant) Platform {
	return DefaultConfigsPerPlatform[dbPlatform.String()]
}

// PlatformDirectory returns the parent directory related to the platform
func (conf Platform) PlatformDirectory() (string, error) {
	basePath, err := conf.appConfig.GetBaseDirectoryPath()

	return path.Join(basePath, conf.directory), err
}

// ArchiveFilepath returns the absolute filepath related to the platform's database archive file
func (conf Platform) ArchiveFilepath() (string, error) {
	platformDirectory, err := conf.PlatformDirectory()

	return path.Join(platformDirectory, conf.source.archiveFilename), err
}

// Filepath returns the absolute filepath related to the  platform's database content file
func (conf Platform) Filepath() (string, error) {
	platformDirectory, err := conf.PlatformDirectory()

	return path.Join(platformDirectory, conf.source.filename), err
}

// IndexFilepath returns absolute path of Index file
func (conf Platform) IndexFilepath() (string, error) {
	platformDirectory, err := conf.PlatformDirectory()

	return path.Join(platformDirectory, conf.index.directory), err
}

// ForceSourceDownload spcifies if the source should be redownloaded, even in the case of source already existing in the filesystem
func (conf Platform) ForceSourceDownload() bool {
	return conf.source.forceDownload
}

// URL returns url associated with the source of the titles databases that should be fetched before parsing
func (conf Platform) URL() string {
	return conf.source.url
}

// IndexVariant returns index type (eg. bleve, solr etc.)
func (conf Platform) IndexVariant() string {
	return conf.index.variant
}

// DocType returns document identifier used for indexes documents matching
func (conf Platform) DocType() string {
	return conf.index.docType
}

// Filename returns filename of source file containing titles database
func (conf Platform) Filename() string {
	return conf.source.filename
}

// PlatformName returns name of the platform whose information is presented in config
func (conf Platform) PlatformName() string {
	return conf.source.name
}

// DatabasePath return path of the database which persists information about entries parsed from the source file
func (conf Platform) DatabasePath() string {
	basePath, err := conf.appConfig.GetBaseDirectoryPath()
	if err != nil {
		return ""
	}

	return path.Join(basePath, conf.database.path)
}

// DatabaseVariant return path of the database which persists information about entries parsed from the source file
func (conf Platform) DatabaseVariant() string {
	return conf.database.variant
}
