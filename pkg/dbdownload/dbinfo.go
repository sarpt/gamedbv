package dbdownload

import (
	"os"
	"path"
)

// DbInfo groups information used for datbase fetching.
type DbInfo struct {
	DbArchiveFileName string
	DbRemoteURL       string
	ForceDbDownload   bool
	LocalDirectory    string
}

type dbFilesStatus struct {
	DoesDatabaseExist bool
}

func downloadDatabaseFile(dbInfo DbInfo) error {
	filePath, err := getDatabaseFilePath(dbInfo)
	if err != nil {
		return err
	}

	outputFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	err = downloadFile(dbInfo.DbRemoteURL, outputFile)
	return err
}

func getFilesStatuses(dbInfo DbInfo) (dbFilesStatus, error) {
	var filesStatus dbFilesStatus

	filePath, err := getDatabaseFilePath(dbInfo)
	if err != nil {
		return filesStatus, err
	}

	filesStatus.DoesDatabaseExist = doesFileExist(filePath)

	return filesStatus, nil
}

func prepareDbDirectory(dbInfo DbInfo) error {
	directory, err := getDatabaseDirectory(dbInfo)
	if err != nil {
		return err
	}

	err = os.MkdirAll(directory, 0700)
	return err
}

func getDatabaseDirectory(dbInfo DbInfo) (string, error) {
	userPath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return path.Join(userPath, databasesBaseDirectories, dbInfo.LocalDirectory), nil
}

func getDatabaseFilePath(dbInfo DbInfo) (string, error) {
	databaseDirectory, err := getDatabaseDirectory(dbInfo)
	if err != nil {
		return "", err
	}

	return path.Join(databaseDirectory, dbInfo.DbArchiveFileName), nil
}
