package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"21BLC1564/routes"
	"21BLC1564/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSearchFiles(t *testing.T) {

	gin.SetMode(gin.TestMode)

	
	utils.ConnectDB()


	router := routes.SetupRouter()

	
	req, _ := http.NewRequest("GET", "/api/search/files?name=testfile", nil)
	
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	
	assert.Equal(t, http.StatusOK, w.Code)

	
	expected := `[]` 
	assert.Contains(t, w.Body.String(), expected)
}
