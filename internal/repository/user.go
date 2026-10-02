package repository

import (
	"errors"
	"fmt"

	"github.com/softserve/go-with-genai-topic4-error-handling/internal/models"
)

type DatabaseError struct {
	Query       string
	OriginalErr error
}

func (e *DatabaseError) Error() string {
	return fmt.Sprintf("db query %q failed: %v", e.Query, e.OriginalErr)
}

func (e *DatabaseError) Unwrap() error {
	return e.OriginalErr
}

type UserRepo struct{}

func (r *UserRepo) FetchUser(id int) (*models.User, error) {
	query := "SELECT id, name FROM users WHERE id = ?"
	if id != 1 {
		return nil, fmt.Errorf("fetch user %d: %w", id, &DatabaseError{
			Query:       query,
			OriginalErr: errors.New("user not found"),
		})
	}
	return &models.User{ID: 1, Name: "leo"}, nil
}
