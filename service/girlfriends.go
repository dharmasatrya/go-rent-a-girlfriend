package service

import (
	"fmt"
	"net/http"
	repository "rent-a-girlfriend/repositories"
)

type GirlfriendService interface {
	GetGirlfriendById(id int) (int, map[string]interface{})
}

type girlfriendService struct {
	girlfriendRepository repository.GirlfriendsRepository
}

func NewGirlfriendService(girlfriendRepository repository.GirlfriendsRepository) *girlfriendService {
	return &girlfriendService{girlfriendRepository}
}

func (s *girlfriendService) GetGirlfriendById(id int) (int, map[string]interface{}) {

	rooms, err := s.girlfriendRepository.GetGirlfriendById(id)
	if err != nil {
		return http.StatusInternalServerError, map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Getting all rooms fail: %v", err),
		}
	}

	return http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Successfully getting girl",
		"data":    rooms,
	}
}
