package repository

import (
	"rent-a-girlfriend/models"

	"github.com/stretchr/testify/mock"
)

type GirlfriendsRepositoryMock struct {
	Mock mock.Mock
}

func (m *GirlfriendsRepositoryMock) GetGirlfriendById(userID int) (*models.Girl, error) {
	res := m.Mock.Called(userID)

	if res.Get(0) == nil {
		return nil, res.Error(1)
	}

	girl := res.Get(0).(*models.Girl)
	return girl, res.Error(1)
}
