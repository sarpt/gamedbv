package dl

import (
	"fmt"
	"os"

	"github.com/sarpt/gamedbv/pkg/config"
	"github.com/sarpt/gamedbv/pkg/platform"
	"github.com/sarpt/gamedbv/pkg/progress"
)

// FilesStatus groups information about existence of specific platform's database files
type FilesStatus struct {
	DoesDatabaseExist bool
}

// DownloadPlatformDatabase downloads neccessary database files related to provided platform
func DownloadPlatformDatabase(variant platform.Variant, printer progress.Notifier) {
	config := config.GetConfig(variant)

	databaseFilesStatuses, err := getFilesStatuses(config)
	if err != nil {
		printer.NextError(err)
		return
	}

	if databaseFilesStatuses.DoesDatabaseExist && !config.ForceSourceDownload() {
		printer.NextProgress(fmt.Sprintf(archiveFileAlreadyPresent, variant.String()))
		return
	}

	err = prepareDatabaseDirectory(config)
	if err != nil {
		printer.NextError(err)
		return
	}

	printer.NextProgress(fmt.Sprintf(downloadingInProgress, config.PlatformName()))
	err = downloadDatabaseFile(config)
	if err != nil {
		printer.NextError(err)
	}
}

func downloadDatabaseFile(config config.Platform) error {
	filePath, err := config.ArchiveFilepath()
	if err != nil {
		return err
	}

	outputFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	err = downloadFile(config.URL(), outputFile)
	return err
}

func getFilesStatuses(config config.Platform) (FilesStatus, error) {
	var filesStatus FilesStatus

	filePath, err := config.ArchiveFilepath()
	if err != nil {
		return filesStatus, err
	}

	filesStatus.DoesDatabaseExist = doesFileExist(filePath)

	return filesStatus, nil
}

func prepareDatabaseDirectory(config config.Platform) error {
	directory, err := config.PlatformDirectory()
	if err != nil {
		return err
	}

	err = os.MkdirAll(directory, 0700)
	return err
}
