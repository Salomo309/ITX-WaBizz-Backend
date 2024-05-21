package handlers

import (
	"io"
	"os"
	"mime/multipart"
	"path/filepath"
)

func SaveFile(fileHeader *multipart.FileHeader) (string, error) {
	// Open the file
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	storageDir := "../storage"

	// Create the storage directory if it does not exist
	if _, err := os.Stat(storageDir); os.IsNotExist(err) {
		err = os.Mkdir(storageDir, os.ModePerm)
		if err != nil {
			return "", err
		}
	}

	// Use the file's original filename
	filename := fileHeader.Filename

	// Define the file path
	filePath := filepath.Join(storageDir, filename)

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

