package models

import (
	"time"
)

// User represents user information
type User struct {
	ID        uint       `json:"id" example:"1"`
	Username  string     `json:"username" example:"johndoe" binding:"required"`
	Password  string     `json:"-" example:"password123" binding:"required"`
	Email     string     `json:"email" example:"john@example.com" binding:"required"`
	Role      string     `json:"role" example:"boy" binding:"required"`
	CreatedAt time.Time  `json:"created_at,omitempty" swaggerignore:"true" example:"2024-01-01T00:00:00Z"`
	UpdatedAt time.Time  `json:"updated_at,omitempty" swaggerignore:"true" example:"2024-01-01T00:00:00Z"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" swaggertype:"string" format:"date-time" example:"2024-01-01T00:00:00Z"`
}

// Boy represents boy profile information
type Boy struct {
	ID                uint       `json:"id" example:"1"`
	UserID            uint       `json:"user_id" example:"1" binding:"required"`
	FirstName         string     `json:"first_name" example:"John" binding:"required"`
	LastName          string     `json:"last_name" example:"Doe" binding:"required"`
	Age               int        `json:"age" example:"25" binding:"required"`
	ProfilePictureURL string     `json:"profile_picture_url" example:"https://example.com/profile.jpg"`
	Bio               string     `json:"bio" example:"I love traveling and meeting new people"`
	User              User       `json:"-" swaggerignore:"true"`
	CreatedAt         time.Time  `json:"created_at,omitempty" swaggerignore:"true" example:"2024-01-01T00:00:00Z"`
	UpdatedAt         time.Time  `json:"updated_at,omitempty" swaggerignore:"true" example:"2024-01-01T00:00:00Z"`
	DeletedAt         *time.Time `json:"deleted_at,omitempty" swaggertype:"string" format:"date-time" example:"2024-01-01T00:00:00Z"`
}

// Girl represents girl profile information
type Girl struct {
	ID                uint       `json:"id" example:"1"`
	UserID            uint       `json:"user_id" example:"1" binding:"required"`
	FirstName         string     `json:"first_name" example:"Jane" binding:"required"`
	LastName          string     `json:"last_name" example:"Doe" binding:"required"`
	Age               int        `json:"age" example:"23" binding:"required"`
	ProfilePictureURL string     `json:"profile_picture_url" example:"https://example.com/profile.jpg"`
	Bio               string     `json:"bio" example:"I enjoy meeting new people"`
	DailyRate         int        `json:"daily_rate" example:"100" binding:"required"`
	User              User       `json:"-" swaggerignore:"true"`
	CreatedAt         time.Time  `json:"created_at,omitempty" swaggerignore:"true" example:"2024-01-01T00:00:00Z"`
	UpdatedAt         time.Time  `json:"updated_at,omitempty" swaggerignore:"true" example:"2024-01-01T00:00:00Z"`
	DeletedAt         *time.Time `json:"deleted_at,omitempty" swaggertype:"string" format:"date-time" example:"2024-01-01T00:00:00Z"`
}

type RegisterRequest struct {
	Username string `json:"username" example:"johndoe" binding:"required"`
	Password string `json:"password" example:"password123" binding:"required"`
	Email    string `json:"email" example:"john@example.com" binding:"required"`
	Role     string `json:"role" example:"boys" binding:"required"`
}

// LoginRequest represents login request information
type LoginRequest struct {
	Email    string `json:"email" example:"john@example.com"`
	Password string `json:"password" example:"password123"`
}

type LoginResponse struct {
	Token string `json:"token" example:"john@example.com"`
}

type LoginRequestWithPassword struct {
	Email    string `json:"email" example:"john@example.com"`
	Password string `json:"password" example:"password123"`
}

// Rating represents rating information
type Rating struct {
	ID        uint       `json:"id" example:"1"`
	GirlID    uint       `json:"girl_id" example:"1" binding:"required"`
	Review    string     `json:"review" example:"Great experience!"`
	Stars     int        `json:"stars" example:"5" binding:"required"`
	Girl      Girl       `json:"-" swaggerignore:"true"`
	CreatedAt time.Time  `json:"created_at,omitempty" swaggerignore:"true" example:"2024-01-01T00:00:00Z"`
	UpdatedAt time.Time  `json:"updated_at,omitempty" swaggerignore:"true" example:"2024-01-01T00:00:00Z"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" swaggertype:"string" format:"date-time" example:"2024-01-01T00:00:00Z"`
}

type CreateBoyProfileRequest struct {
	FirstName         string `json:"first_name" example:"John" binding:"required"`
	LastName          string `json:"last_name" example:"Doe" binding:"required"`
	Age               int    `json:"age" example:"25" binding:"required"`
	ProfilePictureURL string `json:"profile_picture_url" example:"https://example.com/profile.jpg"`
	Bio               string `json:"bio" example:"I love traveling and meeting new people"`
}

type GiveRatingRequest struct {
	GirlID uint   `json:"girl_id" example:"1" binding:"required"`
	Review string `json:"review" example:"Great experience!"`
	Stars  int    `json:"stars" example:"5" binding:"required"`
}

type GirlDetailResponse struct {
	ID                uint     `json:"id"`
	UserID            uint     `json:"user_id"`
	FirstName         string   `json:"first_name"`
	LastName          string   `json:"last_name"`
	Age               int      `json:"age"`
	ProfilePictureURL string   `json:"profile_picture_url"`
	Bio               string   `json:"bio"`
	DailyRate         int      `json:"daily_rate"`
	IsAvailable       bool     `json:"is_available"`
	Ratings           []Rating `json:"ratings"`
}
