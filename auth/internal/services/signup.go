package services

import (
	"auth/internal/core"
	"auth/internal/core/dtos"
	"auth/internal/models"
	"context"
	"errors"
	"log"
)

type signupService struct {
	repo core.UserRepo
}

type SignupServiceConfig struct {
	Repo core.UserRepo
}

func NewSignupService(ssc *SignupServiceConfig) core.UserService {
	return &signupService{
		repo: ssc.Repo,
	}
}

func (ss *signupService) Signup(ctx context.Context, dto *dtos.SignupDto) (*models.User, error) {
	foundUser, err := ss.repo.GetByEmail(ctx, dto.Email)
	if err != nil {
		log.Printf("[Signup Service] | %v \n", err)
	}

	// user is registered before
	if foundUser != nil {
		log.Printf("[Signup Service] | user is already registered before \n")
		return nil, errors.New("user is registered before")
	}

	user, err := ss.repo.Insert(ctx, &models.User{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
		Username:  dto.Username,
		Password:  dto.Password,
	})
	if err != nil {
		log.Printf("[Signup Service] | %v \n", err)
	}

	return user, nil
}
