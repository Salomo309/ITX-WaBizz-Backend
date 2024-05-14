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

	// Read the file's content
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	// Create the storage directory if it does not exist
	storageDir := "storage"
	if _, err := os.Stat(storageDir); os.IsNotExist(err) {
		err = os.Mkdir(storageDir, os.ModePerm)
		if err != nil {
			return "", err
		}
	}

	// Use the file's original filename
	filename := fileHeader.Filename

	// Write the file to the storage directory
	filePath := filepath.Join(storageDir, filename)
	err = os.WriteFile(filePath, fileBytes, 0644)
	if err != nil {
		return "", err
	}

	return filePath, nil
}

