package platform

import (
	"path"

	"github.com/sarpt/gamedbv/pkg/gamedbv"
)

// Config groups information used for platform database handling
type Config struct {
	appConfig         gamedbv.Config
	databaseType      string
	archiveFileName   string
	contentFileName   string
	url               string
	forceDbDownload   bool
	indexType         string
	IndexDir          string
	platformDirectory string
	platformName      string
	docType           string
}

// GetConfig returns final platform database information object taking into accounts default values and passed overrides of settings.
// Todo: implement passing overrides
func GetConfig(dbPlatform Variant) Config {
	return DefaultConfigsPerPlatform[dbPlatform.String()]
}

// PlatformDirectory returns the parent directory related to the platform
func (conf Config) PlatformDirectory() (string, error) {
	homePath, err := conf.appConfig.GetBaseDirectoryPath()

	return path.Join(homePath, conf.platformDirectory), err
}

// PlatformArchiveFilePath returns the absolute filepath related to the platform's database archive file
func (conf Config) PlatformArchiveFilePath() (string, error) {
	databaseDirectory, err := conf.PlatformDirectory()

	return path.Join(databaseDirectory, conf.archiveFileName), err
}

// DatabaseContentFilePath returns the absolute filepath related to the  platform's database content file
func (conf Config) DatabaseContentFilePath() (string, error) {
	databaseDirectory, err := conf.PlatformDirectory()

	return path.Join(databaseDirectory, conf.contentFileName), err
}

// IndexFilePath returns absolute path of Index file
func (conf Config) IndexFilePath() (string, error) {
	databaseDirectory, err := conf.PlatformDirectory()

	return path.Join(databaseDirectory, conf.IndexDir), err
}

// ForceDbDownload spcifies if the source should be redownloaded, even in the case of source already existing in the filesystem
func (conf Config) ForceDbDownload() bool {
	return conf.forceDbDownload
}

// URL returns url associated with the source of the titles databases that should be fetched before parsing
func (conf Config) URL() string {
	return conf.url
}

// IndexType returns index type (eg. bleve, solr etc.)
func (conf Config) IndexType() string {
	return conf.indexType
}

// DocType returns document identifier used for indexes documents matching
func (conf Config) DocType() string {
	return conf.docType
}

// ContentFileName returns filename of source file containing titles database
func (conf Config) ContentFileName() string {
	return conf.contentFileName
}

// Platform returns name of the platform whose information is presented in config
func (conf Config) Platform() string {
	return conf.platformName
}
