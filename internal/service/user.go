package service

import (
	"fmt"

	"github.com/softserve/go-with-genai-topic4-error-handling/internal/models"
	"github.com/softserve/go-with-genai-topic4-error-handling/internal/repository"
)

type UserService struct {
	repo *repository.UserRepo
}

func NewUserService(repo *repository.UserRepo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUser(id int) (*models.User, error) {
	u, err := s.repo.FetchUser(id)
	if err != nil {
		return nil, fmt.Errorf("get user: %w", err)
	}
	return u, nil
}
