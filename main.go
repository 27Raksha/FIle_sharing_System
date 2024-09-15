package main

import (
	"21BLC1564/routes"
	"21BLC1564/utils"
	"log"
)

func main() {
	utils.LoadEnv()

	
	utils.ConnectDB()
	utils.ConnectRedis()

	utils.RunBackgroundJob()
	
	r := routes.SetupRouter()
	r.Static("/uploads", "./uploads")
	log.Fatal(r.Run(":8080"))
}
