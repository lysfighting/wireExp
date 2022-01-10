package service

import (
	"context"
	"wireExp/internal/api/repository"
)

type UserService struct {
	userRepo *repository.UserRepo
}

func NewUserService(userRepo *repository.UserRepo, ) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) Index(ctx context.Context) string {
	return s.userRepo.Exp(ctx)
}
