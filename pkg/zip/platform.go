package zip

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

// PlatformConfig provides information neccessary for unzipping the platform files
type PlatformConfig interface {
	ArchiveFilepath() string
	Filename() string
	Name() string
	Filepath() string
}

// UnzipPlatformDatabase perfoms decompression of platform's database archive file. Returns string with extracted filename, or error
func UnzipPlatformDatabase(config PlatformConfig) error {
	dbArchivePath := config.ArchiveFilepath()
	zipFileReader, err := zip.OpenReader(dbArchivePath)
	if err != nil {
		return err
	}
	defer zipFileReader.Close()

	var contentFileReader io.Reader
	for _, file := range zipFileReader.File {
		if file.Name != config.Filename() {
			continue
		}

		contentFileReader, err = file.Open()
	}

	if contentFileReader == nil {
		return fmt.Errorf(fmt.Sprintf(noDatabaseContentFile, config.Filename(), config.Name()))
	} else if err != nil {
		return err
	}

	contentFilePath := config.Filepath()
	contentFileWriter, err := os.Create(contentFilePath)
	if err != nil {
		return err
	}

	_, err = io.Copy(contentFileWriter, contentFileReader)
	return err
}
