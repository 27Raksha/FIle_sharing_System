package utils

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
)


func SaveFileToLocal(file *multipart.FileHeader) (string, error) {
	uploadPath := "./uploads"
	if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
		err := os.Mkdir(uploadPath, os.ModePerm)
		if err != nil {
			return "", fmt.Errorf("failed to create upload directory: %v", err)
		}
	}

	
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open uploaded file: %v", err)
	}
	defer src.Close()

	
	filePath := filepath.Join(uploadPath, file.Filename)
	dst, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create destination file: %v", err)
	}
	defer dst.Close()

	
	_, err = dst.ReadFrom(src)
	if err != nil {
		return "", fmt.Errorf("failed to save file: %v", err)
	}

	return filePath, nil
}
