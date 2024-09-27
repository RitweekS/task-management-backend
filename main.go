package main

import (
	"log"
	"task-management/database"
	"task-management/routes"

	"github.com/gin-gonic/gin"
)


func main(){
	database.InItDB()
	router := gin.Default()
	routes.InitializeRoutes(router)
	
	if err := router.Run(":8000") ;err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}