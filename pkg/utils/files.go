package utils

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/text/encoding/charmap"
)

func DownloadFile(downloadURL string) ([]byte, error) {
	response, err := http.Get(downloadURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to download, status code: %d", response.StatusCode)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func ExtractJSONFromZip(zipData []byte, filename string) ([]byte, error) {
	reader, err := zip.NewReader(bytes.NewReader(zipData), int64(len(zipData)))
	if err != nil {
		return nil, err
	}

	var jsonData []byte
	for _, file := range reader.File {
		if file.Name == filename {
			rc, err := file.Open()
			if err != nil {
				return nil, err
			}
			defer rc.Close()

			jsonData, err = ioutil.ReadAll(rc)
			if err != nil {
				return nil, err
			}
			break
		}
	}

	if jsonData == nil {
		return nil, fmt.Errorf("*.json not found in the zip file")
	}

	decoder := charmap.ISO8859_1.NewDecoder()
	utf8Data, err := decoder.Bytes(jsonData)
	if err != nil {
		return nil, err
	}

	return utf8Data, nil
}

func downloadFile(downloadURL string) ([]byte, error) {
	response, err := http.Get(downloadURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to download, status code: %d", response.StatusCode)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func extractJSONFromZip(zipData []byte, filename string) ([]byte, error) {
	reader, err := zip.NewReader(bytes.NewReader(zipData), int64(len(zipData)))
	if err != nil {
		return nil, err
	}

	var jsonData []byte
	for _, file := range reader.File {
		if file.Name == filename {
			rc, err := file.Open()
			if err != nil {
				return nil, err
			}
			defer rc.Close()

			jsonData, err = ioutil.ReadAll(rc)
			if err != nil {
				return nil, err
			}
			break
		}
	}

	if jsonData == nil {
		return nil, fmt.Errorf("*.json not found in the zip file")
	}

	decoder := charmap.ISO8859_1.NewDecoder()
	utf8Data, err := decoder.Bytes(jsonData)
	if err != nil {
		return nil, err
	}

	return utf8Data, nil
}

func LoadJSON[T any](jsonData []byte) ([]T, error) {
	var data []T
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
