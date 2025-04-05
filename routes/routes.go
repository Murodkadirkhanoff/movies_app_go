package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	authenticated := server.Group("/")

	// authenticated.Use(middlewares.Authenticate)

	authenticated.POST("/movies", createMovie)
	authenticated.GET("/movies", getMovies)
	authenticated.GET("/movies/:id", getMovieByID)
	authenticated.PUT("/movies/:id", updateMovie)
	authenticated.DELETE("/movies/:id", deleteMovie)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
