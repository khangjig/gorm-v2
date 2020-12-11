package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int64          `json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

type UserResponse struct {
	User *User `json:"user"`
}

type UsersResponse struct {
	Users []User `json:"users"`
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
	Total int64  `json:"total"`
}
