package dl

import (
	"os"

	"github.com/sarpt/gamedbv/internal/config"
	"github.com/sarpt/gamedbv/pkg/platform"
	"github.com/sarpt/gamedbv/pkg/progress"
)

// FilesStatus groups information about existence of specific platform's source files
type FilesStatus struct {
	DoesSourceExist bool
}

// DownloadPlatformSource downloads neccessary source files related to provided platform
func DownloadPlatformSource(appConf config.App, variant platform.Variant, printer progress.Notifier) {
	platformConfig := appConf.Platform(variant)
	sourcesFilesStatuses, err := getFilesStatuses(platformConfig)
	if err != nil {
		printer.NextError(err)
		return
	}

	if sourcesFilesStatuses.DoesSourceExist && !platformConfig.ForceSourceDownload() {
		printer.NextStatus(newArchiveFileAlreadyPresentStatus(variant.String()))
		return
	}

	err = preparePlatformDirectory(platformConfig)
	if err != nil {
		printer.NextError(err)
		return
	}

	printer.NextStatus(newDownloadingInProgressStatus(platformConfig.Name()))
	err = downloadSourceFile(platformConfig)
	if err != nil {
		printer.NextError(err)
	}
}

func downloadSourceFile(config config.Platform) error {
	filePath := config.ArchiveFilepath()
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

	filePath := config.ArchiveFilepath()
	filesStatus.DoesSourceExist = doesFileExist(filePath)

	return filesStatus, nil
}

func preparePlatformDirectory(config config.Platform) error {
	directory := config.DirectoryPath()

	err := os.MkdirAll(directory, 0700)
	return err
}
