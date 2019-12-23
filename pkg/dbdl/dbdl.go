package dbdl

import (
	"fmt"
	"os"

	"github.com/sarpt/gamedbv/pkg/gamedb"
	"github.com/sarpt/gamedbv/pkg/platform"
)

// DownloadPlatformDatabase downloads neccessary database files related to provided platform
func DownloadPlatformDatabase(dbPlatform platform.Variant, progress chan<- string, errors chan<- error) {
	platformDbInfo := gamedb.GetDbInfo(dbPlatform)

	databaseFilesStatuses, err := getFilesStatuses(platformDbInfo)
	if err != nil {
		errors <- err
		return
	}

	if databaseFilesStatuses.DoesDatabaseExist && !platformDbInfo.ForceDbDownload {
		progress <- fmt.Sprintf(archiveFileAlreadyPresent, dbPlatform.String())
		return
	}

	err = prepareDatabaseDirectory(platformDbInfo)
	if err != nil {
		errors <- err
		return
	}

	progress <- fmt.Sprintf(downloadingInProgress, dbPlatform.String())
	err = downloadDatabaseFile(platformDbInfo)
	if err != nil {
		errors <- err
	}
}

func downloadDatabaseFile(platformDbInfo gamedb.Info) error {
	filePath, err := platformDbInfo.GetDatabaseArchiveFilePath()
	if err != nil {
		return err
	}

	outputFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	err = downloadFile(platformDbInfo.URL, outputFile)
	return err
}

func getFilesStatuses(platformDbInfo gamedb.Info) (gamedb.FilesStatus, error) {
	var filesStatus gamedb.FilesStatus

	filePath, err := platformDbInfo.GetDatabaseArchiveFilePath()
	if err != nil {
		return filesStatus, err
	}

	filesStatus.DoesDatabaseExist = doesFileExist(filePath)

	return filesStatus, nil
}

func prepareDatabaseDirectory(platformDbInfo gamedb.Info) error {
	directory, err := platformDbInfo.GetDatabaseDirectory()
	if err != nil {
		return err
	}

	err = os.MkdirAll(directory, 0700)
	return err
}
