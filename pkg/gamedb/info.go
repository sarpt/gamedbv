package gamedb

import (
	"os"
	"path"

	"github.com/sarpt/gamedbv/pkg/platform"
)

// Info groups information used for games database handling
type Info struct {
	ArchiveFileName   string
	ContentFileName   string
	URL               string
	ForceDbDownload   bool
	PlatformDirectory string
}

// FilesStatus groups information about existence of specific platform's database files
type FilesStatus struct {
	DoesDatabaseExist bool
}

const (
	baseDirectory = ".gamedbv"
)

// GetDbInfo returns final platform database information object taking into accounts default values and passed overrides of settings.
// Todo: implement passing overrides
func GetDbInfo(dbPlatform platform.Variant) Info {
	return DefaultDbInfosByPlatform[dbPlatform.String()]
}

// GetDatabaseDirectory returns the parent directory related to the platform
func (ginfo Info) GetDatabaseDirectory() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return path.Join(homePath, baseDirectory, ginfo.PlatformDirectory), nil
}

// GetDatabaseArchiveFilePath returns the absolute filepath related to the platform's database archive file
func (ginfo Info) GetDatabaseArchiveFilePath() (string, error) {
	databaseDirectory, err := ginfo.GetDatabaseDirectory()
	if err != nil {
		return "", err
	}

	return path.Join(databaseDirectory, ginfo.ArchiveFileName), nil
}

// GetDatabaseContentFilePath returns the absolute filepath related to the  platform's database content file
func (ginfo Info) GetDatabaseContentFilePath() (string, error) {
	databaseDirectory, err := ginfo.GetDatabaseDirectory()
	if err != nil {
		return "", err
	}

	return path.Join(databaseDirectory, ginfo.ContentFileName), nil
}
