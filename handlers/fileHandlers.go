package handlers

import (
	"io"
	"os"
	"mime/multipart"
)

func SaveFile(file multipart.File) (error) {
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	err = os.WriteFile("storage/"+ "tes.png", fileBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}