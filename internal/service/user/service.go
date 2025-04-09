package user

import (
	"biliard_club/domain"
	"log/slog"
)

type UserRepository interface {
	GetByPhone(phone string) (*domain.User, error)
	Create(user *domain.User) (*domain.User, error)
}

type Service struct {
	repo UserRepository
}

func NewService(repo UserRepository) *Service {
	return &Service{repo}
}

func (srv *Service) Create(user *domain.User) (*domain.User, error) {
	user, err := srv.repo.Create(user)
	if err != nil {
		slog.Error("failed to create user",
			"error", err.Error())
		return nil, err
	}
	return user, nil
}
