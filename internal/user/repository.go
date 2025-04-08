package user

import (
	"biliard_club/pkg/db"
	"log/slog"
)

type Repository struct {
	Database *db.Db
}

func NewRepository(database *db.Db) *Repository {
	return &Repository{
		Database: database,
	}
}

func (repo *Repository) Create(user *User) (*User, error) {

	result := repo.Database.Create(user)
	if result.Error != nil {
		slog.Error("failed to create user",
			"error", result.Error,
			slog.Group("user",
				"name", user.Name,
				"id", user.ID))
		return nil, result.Error
	}
	slog.Debug("new user created",
		slog.Group("user",
			"name", user.Name,
			"id", user.ID),
	)
	return user, nil
}

func (repo *Repository) GetByID(id uint) (*User, error) {
	var user User
	res := repo.Database.First(&user, "id = ?", id)
	if res.Error != nil {
		slog.Error("failed to get user by id",
			slog.Group(
				"user",
				"id", id))
		return nil, res.Error
	}
	return &user, nil
}

func (repo *Repository) GetByPhone(phone string) (*User, error) {
	var user User
	res := repo.Database.First(&user, "phone = ?", phone)
	if res.Error != nil {
		slog.Error("failed to get user by phone",
			slog.Group(
				"user",
				"phone", phone))
		return nil, res.Error
	}
	return &user, nil
}
