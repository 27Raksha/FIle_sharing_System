package controllers_test

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"21BLC1564/routes"
	"21BLC1564/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUploadFile(t *testing.T) {
	
	gin.SetMode(gin.TestMode)

	
	utils.ConnectDB()


	router := routes.SetupRouter()


	filePath := "testfile.txt"
	file, err := os.Create(filePath)
	assert.NoError(t, err)  
	defer file.Close()      

	
	_, err = file.WriteString("This is a test file")
	assert.NoError(t, err)   


	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filePath)
	assert.NoError(t, err)  


	testFile, err := os.Open(filePath)
	assert.NoError(t, err)   
	defer testFile.Close()   

	_, err = io.Copy(part, testFile)
	assert.NoError(t, err)   
	writer.Close()   
	req, _ := http.NewRequest("POST", "/api/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	
	w := httptest.NewRecorder()

	
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected status code 200, but got %v", w.Code)
	err = os.Remove(filePath)
	assert.NoError(t, err, "Failed to remove test file")
}
