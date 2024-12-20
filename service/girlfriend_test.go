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
	girlfriendRepoMock.Mock.On("GetGirlfriendById", 1).Return(&mockGirl, nil).Once()

	// Call the service method
	status, response := testGirlfriendService.GetGirlfriendById(int(mockGirl.ID))

	// Assertions
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, "Successfully getting girl", response["message"])
	assert.NotNil(t, response["data"])
	assert.Equal(t, &mockGirl, response["data"])
}

func TestCreateRating(t *testing.T) {
	// Mock the response for successful rating creation
	mockRating := models.Rating{
		GirlID:    1,
		Review:    "aaaaaa",
		Stars:     3,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Mock the CreateRating call to return the mockRating
	girlfriendRepoMock.Mock.On("CreateRating", &mockRating).Return(&mockRating, nil).Once()

	// Call the service method
	status, response := testGirlfriendService.CreateRating(&mockRating)

	// Assertions
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, "Successfully created rating", response["message"])
	assert.NotNil(t, response["data"])

	// Compare the pointers directly
	assert.Equal(t, mockRating, response["data"])
}
