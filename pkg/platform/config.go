package platform

import (
	"path"

	"github.com/sarpt/gamedbv/pkg/gamedbv"
)

// Config groups information used for platform database handling
type Config struct {
	AppConfig         gamedbv.Config
	DatabaseType      string
	ArchiveFileName   string
	ContentFileName   string
	URL               string
	ForceDbDownload   bool
	IndexType         string
	IndexDir          string
	PlatformDirectory string
	PlatformName      string
}

// GetConfig returns final platform database information object taking into accounts default values and passed overrides of settings.
// Todo: implement passing overrides
func GetConfig(dbPlatform Variant) Config {
	return DefaultConfigsPerPlatform[dbPlatform.String()]
}

// GetPlatformDirectory returns the parent directory related to the platform
func (conf Config) GetPlatformDirectory() (string, error) {
	homePath, err := conf.AppConfig.GetBaseDirectoryPath()

	return path.Join(homePath, conf.PlatformDirectory), err
}

// GetPlatformArchiveFilePath returns the absolute filepath related to the platform's database archive file
func (conf Config) GetPlatformArchiveFilePath() (string, error) {
	databaseDirectory, err := conf.GetPlatformDirectory()

	return path.Join(databaseDirectory, conf.ArchiveFileName), err
}

// GetDatabaseContentFilePath returns the absolute filepath related to the  platform's database content file
func (conf Config) GetDatabaseContentFilePath() (string, error) {
	databaseDirectory, err := conf.GetPlatformDirectory()

	return path.Join(databaseDirectory, conf.ContentFileName), err
}

// GetIndexFilePath returns absolute path of Index file
func (conf Config) GetIndexFilePath() (string, error) {
	databaseDirectory, err := conf.GetPlatformDirectory()

	return path.Join(databaseDirectory, conf.IndexDir), err
}

// GetIndexType returns index type (eg. bleve, solr etc.)
func (conf Config) GetIndexType() string {
	return conf.IndexType
}

// GetPlatform returns platform name
func (conf Config) GetPlatform() string {
	return conf.PlatformName
}
