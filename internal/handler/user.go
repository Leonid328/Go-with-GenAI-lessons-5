package handler

import (
	"fmt"

	"github.com/softserve/go-with-genai-topic4-error-handling/internal/service"
)

type UserHandler struct {
	svc *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) ShowUser(id int) error {
	u, err := h.svc.GetUser(id)
	if err != nil {
		return fmt.Errorf("show user: %w", err)
	}
	fmt.Printf("user: %d %s\n", u.ID, u.Name)
	return nil
}
