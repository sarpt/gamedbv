package dl

import (
	"os"

	"github.com/sarpt/gamedbv/pkg/platform"
	"github.com/sarpt/gamedbv/pkg/progress"
)

// FilesStatus groups information about existence of specific platform's source files
type FilesStatus struct {
	DoesSourceExist bool
}

// Config instructs Dl how to download source file and where to put it
type Config struct {
	URL             string
	Filepath        string
	DirectoryPath   string
	ForceRedownload bool
	PlatformName    string
}

// DownloadPlatformSource downloads neccessary source files related to provided platform
func DownloadPlatformSource(conf Config, variant platform.Variant, printer progress.Notifier) {
	sourcesFilesStatuses, err := getFilesStatuses(conf)
	if err != nil {
		printer.NextError(err)
		return
	}

	if sourcesFilesStatuses.DoesSourceExist && !conf.ForceRedownload {
		printer.NextStatus(newArchiveFileAlreadyPresentStatus(variant.String()))
		return
	}

	err = preparePlatformDirectory(conf)
	if err != nil {
		printer.NextError(err)
		return
	}

	printer.NextStatus(newDownloadingInProgressStatus(conf.PlatformName))
	err = downloadSourceFile(conf)
	if err != nil {
		printer.NextError(err)
	}
}

func downloadSourceFile(conf Config) error {
	filePath := conf.Filepath
	outputFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	err = downloadFile(conf.URL, outputFile)
	return err
}

func getFilesStatuses(conf Config) (FilesStatus, error) {
	var filesStatus FilesStatus

	filePath := conf.Filepath
	filesStatus.DoesSourceExist = doesFileExist(filePath)

	return filesStatus, nil
}

func preparePlatformDirectory(conf Config) error {
	directory := conf.DirectoryPath

	err := os.MkdirAll(directory, 0700)
	return err
}
