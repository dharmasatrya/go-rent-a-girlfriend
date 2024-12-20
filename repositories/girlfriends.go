package repository

import (
	"rent-a-girlfriend/models"

	"gorm.io/gorm"
)

type GirlfriendsRepository interface {
	GetGirlfriendById(userID int) (*models.Girl, error)
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
