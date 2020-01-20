package zip

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

// Config provides information neccessary for unzipping the platform files
type Config interface {
	PlatformArchiveFilePath() (string, error)
	ContentFileName() string
	Platform() string
	DatabaseContentFilePath() (string, error)
}

// UnzipPlatformDatabase perfoms decompression of platform's database archive file. Returns string with extracted filename, or error
func UnzipPlatformDatabase(config Config) error {
	dbArchivePath, err := config.PlatformArchiveFilePath()
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
		if file.Name != config.ContentFileName() {
			continue
		}

		contentFileReader, err = file.Open()
	}

	if contentFileReader == nil {
		return fmt.Errorf(fmt.Sprintf(noDatabaseContentFile, config.ContentFileName(), config.Platform()))
	} else if err != nil {
		return err
	}

	contentFilePath, err := config.DatabaseContentFilePath()
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
