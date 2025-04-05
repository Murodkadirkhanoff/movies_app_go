package main

import (
	"log"

	"mk/movie-app/db/db"
	"mk/movie-app/db/seeder"
	"mk/movie-app/models"
	"mk/movie-app/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	db.InitDB()
	seeder.SeedDatabase()
	db.DB.AutoMigrate(&models.Plot{}, &models.Movie{}, &models.User{})

	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
