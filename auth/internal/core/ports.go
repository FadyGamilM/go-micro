package core

import (
	"auth/internal/core/dtos"
	"auth/internal/models"
	"context"
)

type UserRepo interface {
	Insert(ctx context.Context, user *models.User) (*models.User, error)
	GetByID(ctx context.Context, id int64) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	Update(ctx context.Context, user *models.User) (*models.User, error)
	DeleteByID(ctx context.Context, id int64) error
	DeleteByEmail(ctx context.Context, email string) error
}

type UserService interface {
	Signup(ctx context.Context, dto *dtos.SignupDto) (*models.User, error)
	// Login(ctx context.Context, dto *dtos.LoginDto) (*models.User, error)
}
