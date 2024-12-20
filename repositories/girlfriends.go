package repository

import (
	"rent-a-girlfriend/models"

	"gorm.io/gorm"
)

type GirlfriendsRepository interface {
	GetGirlfriendById(userID int) (*models.Girl, error)
	CreateRating(rating *models.Rating) (*models.Rating, error)
}

type girlfriendsRepository struct {
	db *gorm.DB
}

func NewGirlfriendsRepository(db *gorm.DB) *girlfriendsRepository {
	return &girlfriendsRepository{db}
}

func (r *girlfriendsRepository) GetGirlfriendById(userID int) (*models.Girl, error) {
	var girl models.Girl

	if err := r.db.Where("user_id = ?", userID).Find(&girl).Error; err != nil {
		return nil, err
	}

	return &girl, nil
}

func (r *girlfriendsRepository) CreateRating(rating *models.Rating) (*models.Rating, error) {

	if err := r.db.Create(&rating).Error; err != nil {
		return nil, err
	}

	return rating, nil
}
