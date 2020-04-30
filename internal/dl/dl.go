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
func DownloadPlatformSource(cfg Config, variant platform.Variant, printer progress.Notifier) {
	sourcesFilesStatuses, err := getFilesStatuses(cfg)
	if err != nil {
		printer.NextError(err)
		return
	}

	if sourcesFilesStatuses.DoesSourceExist && !cfg.ForceRedownload {
		printer.NextStatus(newArchiveFileAlreadyPresentStatus(variant.String()))
		return
	}

	err = preparePlatformDirectory(cfg)
	if err != nil {
		printer.NextError(err)
		return
	}

	printer.NextStatus(newDownloadingInProgressStatus(cfg.PlatformName))
	err = downloadSourceFile(cfg)
	if err != nil {
		printer.NextError(err)
	}
}

func downloadSourceFile(cfg Config) error {
	filePath := cfg.Filepath
	outputFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	err = downloadFile(cfg.URL, outputFile)
	return err
}

func getFilesStatuses(cfg Config) (FilesStatus, error) {
	var filesStatus FilesStatus

	filePath := cfg.Filepath
	filesStatus.DoesSourceExist = doesFileExist(filePath)

	return filesStatus, nil
}

func preparePlatformDirectory(cfg Config) error {
	directory := cfg.DirectoryPath

	err := os.MkdirAll(directory, 0700)
	return err
}
