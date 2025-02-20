package handlers

import (
	"io"
	"itx-wabizz/configs"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"
)

func saveFile(fileHeader *multipart.FileHeader) (string, error) {
	// Open the file
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Create the storage directory if it does not exist
	if _, err := os.Stat(configs.StoragePath); os.IsNotExist(err) {
		err = os.Mkdir(configs.StoragePath, os.ModePerm)
		if err != nil {
			return "", err
		}
	}

	// Use the file's original filename
	filename := fileHeader.Filename

	// Define the file path
	filePath := filepath.Join(configs.StoragePath, filename)

	// Create a new file in the storage directory
	outFile, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer outFile.Close()

	// Copy the uploaded file's content to the new file
	_, err = io.Copy(outFile, file)
	if err != nil {
		return "", err
	}

	return filePath, nil
}

func downloadFile(url string) (*multipart.FileHeader, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	tempFile, err := os.CreateTemp("", "downloaded-")
	if err != nil {
		return nil, err
	}
	defer tempFile.Close()

	_, err = io.Copy(tempFile, resp.Body)
	if err != nil {
		return nil, err
	}

	fileInfo, err := tempFile.Stat()
	if err != nil {
		return nil, err
	}

	return &multipart.FileHeader{
		Filename: fileInfo.Name(),
		Size:     fileInfo.Size(),
		Header:   textproto.MIMEHeader{},
	}, nil
}
