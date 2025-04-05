package routes

import (
	"mk/movie-app/db/db"
	"mk/movie-app/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func createMovie(c *gin.Context) {
	var movie models.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx := db.DB.Begin()
	if err := tx.Create(&movie).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := tx.Preload("Plot").First(&movie, movie.ID).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movie with plot"})
		return
	}

	tx.Commit()
	c.JSON(http.StatusCreated, movie)
}

func getMovies(c *gin.Context) {
	var movies []models.Movie
	if err := db.DB.Preload("Plot").Find(&movies).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, movies)
}

func getMovieByID(c *gin.Context) {
	id := c.Param("id")
	var movie models.Movie
	if err := db.DB.Preload("Plot").First(&movie, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}
	c.JSON(http.StatusOK, movie)
}

func updateMovie(c *gin.Context) {
	id := c.Param("id")
	var movie models.Movie
	if err := db.DB.First(&movie, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx := db.DB.Begin()
	if err := tx.Save(&movie).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := tx.Preload("Plot").First(&movie, movie.ID).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movie with plot"})
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, movie)
}

func deleteMovie(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := db.DB.Delete(&models.Movie{}, idUint).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Movie deleted"})
}
