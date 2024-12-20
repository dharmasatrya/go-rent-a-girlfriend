package service_test

import (
	"net/http"
	"rent-a-girlfriend/models"
	repository "rent-a-girlfriend/repositories"
	"rent-a-girlfriend/service"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var girlfriendRepoMock = &repository.GirlfriendsRepositoryMock{Mock: mock.Mock{}}

// Service under test
var testGirlfriendService = service.NewGirlfriendService(girlfriendRepoMock)

func TestGetGirlfriendById(t *testing.T) {
	// Mock the response for successful room retrieval
	mockGirl := models.Girl{
		ID:                1,
		UserID:            1,
		FirstName:         "lisa",
		LastName:          "kim",
		Age:               23,
		ProfilePictureURL: "aaaa",
		Bio:               "aaaaaaa",
		DailyRate:         760000,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	// Mock the GetAllRooms call to return the mockRooms
	girlfriendRepoMock.Mock.On("GetGirlfriendById").Return(&mockGirl, nil).Once()

	// Call the service method
	status, response := testGirlfriendService.GetGirlfriendById(int(mockGirl.ID))

	// Assertions
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, "Successfully getting girl", response["message"])
	assert.NotNil(t, response["data"])
	assert.Equal(t, &mockGirl, response["data"])
}
