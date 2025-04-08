package user

import (
	"log/slog"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo}
}
func (srv *Service) Create(user *User) (*User, error) {
	user, err := srv.repo.Create(user)
	if err != nil {
		slog.Error("failed to create user",
			"error", err.Error())
		return nil, err
	}
	return user, nil
}
