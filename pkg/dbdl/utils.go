package dbdl

import (
	"io"
	"net/http"
	"os"
)

func crateFileGetRequest(url string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "gamedbv")

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
