package dbdownload

import (
	"io"
	"net/http"
	"os"
)

const (
	databasesBaseDirectories = ".gamedbv"
)

// DownloadPlatformDatabase downloads neccessary files related to provided platform
func DownloadPlatformDatabase(platform Platform, messageChannel chan<- string, errorChannel chan<- error) {
	dbInfo := getDbInfo(platform)

	databaseFilesStatuses, err := getFilesStatuses(dbInfo)
	if err != nil {
		errorChannel <- err
		return
	}

	if databaseFilesStatuses.DoesDatabaseExist && !dbInfo.ForceDbDownload {
		messageChannel <- platform.String() + " database will not be downloaded since its already present"
		return
	}

	err = prepareDbDirectory(dbInfo)
	if err != nil {
		errorChannel <- err
		return
	}

	messageChannel <- "Downloading platform database variant of " + platform.String()
	err = downloadDatabaseFile(dbInfo)
	if err != nil {
		errorChannel <- err
	}
}

func getDbInfo(platform Platform) DbInfo {
	return DefaultDbInfosByPlatform[platform.value]
}

func downloadFile(url string, outputFile *os.File) error {
	fileGetRequest, err := crateFileGetRequest(url)
	if err != nil {
		return err
	}

	client := &http.Client{}
	fileGetResponse, err := client.Do(fileGetRequest)
	if err != nil {
		return err
	}
	defer fileGetResponse.Body.Close()

	_, err = io.Copy(outputFile, fileGetResponse.Body)
	return err
}

func crateFileGetRequest(url string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Gamedbf")

	return req, nil
}

func doesFileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	} else if err != nil {
		return false
	}

	return true
}
