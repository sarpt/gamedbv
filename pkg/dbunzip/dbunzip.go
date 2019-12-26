package dbunzip

import (
	"archive/zip"
	"fmt"
	"io"
	"os"

	"github.com/sarpt/gamedbv/pkg/gamedb"
	"github.com/sarpt/gamedbv/pkg/platform"
	"github.com/sarpt/gamedbv/pkg/progress"
)

// UnzipPlatformDatabase perfoms decompression of platform's database archive file
func UnzipPlatformDatabase(platformDb platform.Variant, printer progress.Notifier) {
	dbInfo := gamedb.GetDbInfo(platformDb)

	dbArchivePath, err := dbInfo.GetDatabaseArchiveFilePath()
	if err != nil {
		printer.NextError(err)
		return
	}

	zipFileReader, err := zip.OpenReader(dbArchivePath)
	if err != nil {
		printer.NextError(err)
		return
	}
	defer zipFileReader.Close()

	var contentFileReader io.Reader
	for _, file := range zipFileReader.File {
		if file.Name == dbInfo.ContentFileName {
			contentFileReader, err = file.Open()
		}
	}

	if contentFileReader == nil {
		printer.NextProgress(fmt.Sprintf(noDatabaseContentFile, dbInfo.ContentFileName, platformDb.String()))
		return
	} else if err != nil {
		printer.NextError(err)
		return
	}

	contentFilePath, err := dbInfo.GetDatabaseContentFilePath()
	if err != nil {
		printer.NextError(err)
		return
	}

	contentFileWriter, err := os.Create(contentFilePath)
	if err != nil {
		printer.NextError(err)
		return
	}

	printer.NextProgress(fmt.Sprintf(extractingInProgress, dbInfo.ContentFileName, platformDb.String()))
	_, err = io.Copy(contentFileWriter, contentFileReader)
	if err != nil {
		printer.NextError(err)
	}
}
