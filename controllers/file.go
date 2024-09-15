package controllers

import (
	"time"
	"21BLC1564/models"
	"21BLC1564/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"sync"
)


func UploadFile(c *gin.Context) {
	email, _ := c.Get("email") 

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		
		filePath, err := utils.SaveFileToLocal(file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		
		var user models.User
		if err := utils.DB.Where("email = ?", email).First(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
			return
		}

		
		uploadedFile := models.File{
			UserID:    user.ID,
			Name:      filepath.Base(file.Filename),
			Location:  filePath,
			UploadedAt: time.Now(),
			Size:      file.Size,
		}

		utils.DB.Create(&uploadedFile)

		
		utils.InvalidateCachedFiles(user.ID)
	}()

	
	wg.Wait()

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "file_path": file.Filename})
}


func ListFiles(c *gin.Context) {
	email, _ := c.Get("email") 

	
	var user models.User
	if err := utils.DB.Where("email = ?", email).First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	
	cachedFiles, err := utils.GetCachedFiles(user.ID)
	if err == nil && cachedFiles != nil {
		
		c.JSON(http.StatusOK, cachedFiles)
		return
	}

	
	var files []models.File
	if err := utils.DB.Where("user_id = ?", user.ID).Find(&files).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve files"})
		return
	}

	
	if err := utils.CacheFiles(user.ID, files); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cache files"})
		return
	}

	c.JSON(http.StatusOK, files)
}


func ShareFile(c *gin.Context) {
	fileID := c.Param("file_id") 

	
	var file models.File
	if err := utils.DB.First(&file, fileID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	
	sharedLink := "http://localhost:8080/uploads/" + file.Location

	
	err := utils.CacheSharedLink(fileID, sharedLink, 24*time.Hour) // Cache for 24 hours
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cache shared link"})
		return
	}

	
	c.JSON(http.StatusOK, gin.H{"shared_link": sharedLink})
}


func SearchFiles(c *gin.Context) {
	email, _ := c.Get("email") 


	name := c.Query("name")
	uploadDate := c.Query("uploaded_at") 
	fileType := c.Query("file_type")

	
	var user models.User
	if err := utils.DB.Where("email = ?", email).First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	
	var files []models.File
	query := utils.DB.Where("user_id = ?", user.ID) 

	
	if name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}

	if uploadDate != "" {
		query = query.Where("DATE(uploaded_at) = ?", uploadDate) 
	}

	
	if fileType != "" {
		query = query.Where("name ILIKE ?", "%."+fileType)
	}


	if err := query.Find(&files).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to search files"})
		return
	}

	
	c.JSON(http.StatusOK, files)
}