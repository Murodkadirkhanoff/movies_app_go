package seeder

import (
	"fmt"
	"log"

	"mk/movie-app/db/db"
	"mk/movie-app/models"
	"mk/movie-app/utils"
	"time"
)

func SeedDatabase() {

	seedPlots()
	seedMovies()
	seedUsers()
}

func seedPlots() {
	// Пытаемся создать тестовые данные для Plot
	plots := []models.Plot{
		{Title: "Plot 1", CreatedAt: time.Now()},
		{Title: "Plot 2", CreatedAt: time.Now()},
		{Title: "Plot 3", CreatedAt: time.Now()},
	}

	for _, plot := range plots {
		if err := db.DB.Create(&plot).Error; err != nil {
			log.Printf("Ошибка при создании записи Plot: %v", err)
		} else {
			fmt.Println("Plot добавлен:", plot.Title)
		}
	}
}

func seedMovies() {
	// Пытаемся создать тестовые данные для Movie
	movies := []models.Movie{
		{Title: "Movie 1", Director: "Director 1", Year: 2000, PlotID: nil},
		{Title: "Movie 2", Director: "Director 2", Year: 2005, PlotID: nil},
		{Title: "Movie 3", Director: "Director 3", Year: 2010, PlotID: nil},
	}

	for _, movie := range movies {
		if err := db.DB.Create(&movie).Error; err != nil {
			log.Printf("Ошибка при создании записи Movie: %v", err)
		} else {
			fmt.Println("Movie добавлен:", movie.Title)
		}
	}
}

func seedUsers() {
	// Пытаемся создать тестовые данные для User
	users := []models.User{
		{Email: "user1@example.com", Password: "password123"},
		{Email: "user2@example.com", Password: "password456"},
		{Email: "user3@example.com", Password: "password789"},
	}

	for _, user := range users {
		// Хеширование пароля перед сохранением
		hashedPassword, err := utils.HashPassword(user.Password)
		if err != nil {
			log.Printf("Ошибка хеширования пароля: %v", err)
			continue
		}
		user.Password = hashedPassword

		if err := db.DB.Create(&user).Error; err != nil {
			log.Printf("Ошибка при создании записи User: %v", err)
		} else {
			fmt.Println("User добавлен:", user.Email)
		}
	}
}
