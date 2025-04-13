package user

import (
	"biliard_club/domain"
	"biliard_club/domain/models"
	"errors"
	"log/slog"
)

type Service struct {
	repo domain.UserRepository
}

func NewService(repo domain.UserRepository) *Service {
	return &Service{repo}
}

func (srv *Service) Create(user *models.User) (*models.User, error) {
	user, err := srv.repo.Create(user)
	if err != nil && !errors.Is(err, domain.ErrConflict) {
		slog.Error("failed to create user",
			"error", err.Error())
		return nil, domain.ErrInternalServer
	}
	return user, err
}
