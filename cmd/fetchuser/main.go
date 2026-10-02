package main

import (
	"errors"
	"fmt"

	"github.com/softserve/go-with-genai-topic4-error-handling/internal/handler"
	"github.com/softserve/go-with-genai-topic4-error-handling/internal/repository"
	"github.com/softserve/go-with-genai-topic4-error-handling/internal/service"
)

func main() {
	repo := &repository.UserRepo{}
	svc := service.NewUserService(repo)
	h := handler.NewUserHandler(svc)

	err := h.ShowUser(7)
	if err != nil {
		var dbErr *repository.DatabaseError
		if errors.As(err, &dbErr) {
			fmt.Println("failed query:", dbErr.Query)
			fmt.Println("reason:", dbErr.OriginalErr)
			return
		}
		fmt.Println("error:", err)
	}
}
