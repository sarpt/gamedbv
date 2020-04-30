package zip

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

// Config provides information neccessary for unzipping the platform files
type Config struct {
	ArchiveFilepath string
	SourceFilename  string
	Name            string
	OutputFilepath  string
}

// UnzipPlatformSource perfoms decompression of platform's source archive file. Returns string with extracted filename, or error
func UnzipPlatformSource(config Config) error {
	dbArchivePath := config.ArchiveFilepath
	zipFileReader, err := zip.OpenReader(dbArchivePath)
	if err != nil {
		return err
	}
	defer zipFileReader.Close()

	var contentFileReader io.Reader
	for _, file := range zipFileReader.File {
		if file.Name != config.SourceFilename {
			continue
		}

		contentFileReader, err = file.Open()
	}

	if contentFileReader == nil {
		return fmt.Errorf(fmt.Sprintf(noDatabaseContentFile, config.SourceFilename, config.Name))
	} else if err != nil {
		return err
	}

	contentFilePath := config.OutputFilepath
	contentFileWriter, err := os.Create(contentFilePath)
	if err != nil {
		return err
	}

	_, err = io.Copy(contentFileWriter, contentFileReader)
	return err
}
