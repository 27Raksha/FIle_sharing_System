package utils

import (
	"log"
	"os"
	"time"
	"21BLC1564/models"
)


func RunBackgroundJob() {
	ticker := time.NewTicker(24 * time.Hour) 

	go func() {
		
		for range ticker.C {
			
			DeleteExpiredFiles()
		}
	}()
}


func DeleteExpiredFiles() {
	
	threshold := time.Now().AddDate(0, 0, -30)

	var files []models.File
	
	err := DB.Where("uploaded_at < ?", threshold).Find(&files).Error
	if err != nil {
		log.Println("Error fetching expired files:", err)
		return
	}

	
	for _, file := range files {
		err := DeleteFileFromStorage(file.Location)
		if err != nil {
			log.Println("Error deleting file from storage:", err)
			continue
		}

		
		DB.Delete(&file)
		log.Println("Deleted file:", file.Name)
	}
}


func DeleteFileFromStorage(location string) error {
	err := os.Remove(location) 
	if err != nil {
		return err
	}
	
	return nil
}
