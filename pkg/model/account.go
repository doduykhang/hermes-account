package model

import "time"

type Account struct {
	ID string 
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Email string `gorm:"unique"`
	Password string 
	UserID string `gorm:"unique"`
}
