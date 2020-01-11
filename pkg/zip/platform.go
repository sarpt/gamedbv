package zip

import (
	"archive/zip"
	"fmt"
	"io"
	"os"

	"github.com/sarpt/gamedbv/pkg/platform"
)

// UnzipPlatformDatabase perfoms decompression of platform's database archive file. Returns string with extracted filename, or error
func UnzipPlatformDatabase(variant platform.Variant) error {
	config := platform.GetConfig(variant)

	dbArchivePath, err := config.GetPlatformArchiveFilePath()
	if err != nil {
		return err
	}

	zipFileReader, err := zip.OpenReader(dbArchivePath)
	if err != nil {
		return err
	}
	defer zipFileReader.Close()

	var contentFileReader io.Reader
	for _, file := range zipFileReader.File {
		if file.Name != config.ContentFileName {
			continue
		}

		contentFileReader, err = file.Open()
	}

	if contentFileReader == nil {
		return fmt.Errorf(fmt.Sprintf(noDatabaseContentFile, config.ContentFileName, variant.String()))
	} else if err != nil {
		return err
	}

	contentFilePath, err := config.GetDatabaseContentFilePath()
	if err != nil {
		return err
	}

	contentFileWriter, err := os.Create(contentFilePath)
	if err != nil {
		return err
	}

	_, err = io.Copy(contentFileWriter, contentFileReader)
	return err
}
