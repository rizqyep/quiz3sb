package main

import (
	"log"
	"quiz3-rizqyep/database"
	"quiz3-rizqyep/routers"
	"quiz3-rizqyep/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	envMap := utils.InitEnv()

	DB := database.GetDBConnection()
	database.DBMigrate()
	defer DB.Close()

	r := gin.Default()
	routers.RouteHandlers(r)

	if err := r.Run(":" + envMap["PORT"]); err != nil {
		log.Fatalf("Error in Starting the HTTP Server, Err: %s", err.Error())
	}

}
