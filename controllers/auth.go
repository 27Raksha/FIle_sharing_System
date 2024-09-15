package controllers

import (
    "21BLC1564/models"
    "21BLC1564/utils"
    "github.com/gin-gonic/gin"
    "net/http"
)

func Register(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user.HashPassword(user.Password)
    if result := utils.DB.Create(&user); result.Error != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "could not register user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "user registered successfully"})
}

func Login(c *gin.Context) {
    var user models.User
    var input models.User

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    utils.DB.Where("email = ?", input.Email).First(&user)
    if err := user.CheckPassword(input.Password); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
        return
    }

    token, _ := utils.GenerateJWT(user.Email)
    c.JSON(http.StatusOK, gin.H{"token": token})
}
