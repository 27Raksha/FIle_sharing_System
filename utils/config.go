package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"21BLC1564/models"
	"github.com/gin-gonic/gin"
)

var DB *gorm.DB
var RedisClient *redis.Client


func LoadEnv() {
	
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}


func GetEnv(key string) string {
	return os.Getenv(key)
}


func ConnectDB() {
	LoadEnv() 

	
	var dsn string
	if gin.Mode() == gin.TestMode {
		dsn = os.Getenv("DATABASE_TEST_URL") 
	} else {
		dsn = os.Getenv("DATABASE_URL")
	}

	
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	
	err = DB.AutoMigrate(&models.User{}, &models.File{})
	if err != nil {
		log.Fatal("Failed to migrate the database:", err)
	}

	log.Println("Database connection and migration successful!")
}


func ConnectRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     GetEnv("REDIS_ADDR"),
		Password: GetEnv("REDIS_PASSWORD"), 
		DB:       0,                        
	})

	
	_, err := RedisClient.Ping(RedisClient.Context()).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}
	fmt.Println("Redis connection successful!")
}
