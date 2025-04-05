package models

import (
	"errors"
	"mk/movie-app/db/db"
	"mk/movie-app/utils"
	"time"
)

type Plot struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string    `gorm:"unique;not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type Movie struct {
	ID       uint   `gorm:"primaryKey"`
	Title    string `gorm:"unique;not null" json:"title" binding:"required"`
	Director string `gorm:"not null" json:"director" binding:"required"`
	Year     int    `json:"year" binding:"required,min=1888,max=2100"` // логичный диапазон лет
	PlotID   *uint  `json:"plot_id"`
	Plot     Plot   `gorm:"constraint:OnDelete:SET NULL;" json:"plot,omitempty"`
}

type User struct {
	ID       int64  `gorm:"primaryKey"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

func (u *User) Save() error {
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hashedPassword
	return db.DB.Create(u).Error
}

func (u *User) ValidateCredentials() error {
	var user User
	if err := db.DB.Where("email = ?", u.Email).First(&user).Error; err != nil {
		return errors.New("Credentials invalid")
	}

	if !utils.CheckPasswordHash(u.Password, user.Password) {
		return errors.New("Credentials invalid")
	}

	u.ID = user.ID
	return nil
}
