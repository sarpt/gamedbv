package dbunzip

import (
	"archive/zip"
	"fmt"
	"io"
	"os"

	"github.com/sarpt/gamedbv/pkg/gamedb"
	"github.com/sarpt/gamedbv/pkg/platform"
)

// UnzipPlatformDatabase perfoms decopresion of platform's database archive file
func UnzipPlatformDatabase(platformDb platform.Variant, progress chan<- string, errors chan<- error) {
	dbInfo := gamedb.GetDbInfo(platformDb)

	dbArchivePath, err := dbInfo.GetDatabaseArchiveFilePath()
	if err != nil {
		errors <- err
		return
	}

	zipFileReader, err := zip.OpenReader(dbArchivePath)
	if err != nil {
		errors <- err
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
		progress <- fmt.Sprintf(noDatabaseContentFile, dbInfo.ContentFileName, platformDb.String())
		return
	} else if err != nil {
		errors <- err
		return
	}

	contentFilePath, err := dbInfo.GetDatabaseContentFilePath()
	if err != nil {
		errors <- err
		return
	}

	contentFileWriter, err := os.Create(contentFilePath)
	if err != nil {
		errors <- err
		return
	}

	progress <- fmt.Sprintf(extractingInProgress, dbInfo.ContentFileName, platformDb.String())
	_, err = io.Copy(contentFileWriter, contentFileReader)
	if err != nil {
		errors <- err
	}
}
