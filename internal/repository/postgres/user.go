package postgres

import (
	"biliard_club/domain"
	"biliard_club/pkg/db"
	"errors"
	"gorm.io/gorm"
)

type UserRepository struct {
	Database *db.Db
}

func NewUserRepository(database *db.Db) *UserRepository {
	return &UserRepository{
		Database: database,
	}
}

func (repo *UserRepository) Create(user *domain.User) (*domain.User, error) {
	result := repo.Database.Create(user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, domain.ErrUserExists
		}
		return nil, result.Error
	}
	return user, nil
}

func (repo *UserRepository) GetByID(id uint) (*domain.User, error) {
	var user domain.User
	res := repo.Database.First(&user, "id = ?", id)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, domain.ErrUserNotFound
		}
		return nil, res.Error
	}
	return &user, nil
}

func (repo *UserRepository) GetByPhone(phone string) (*domain.User, error) {
	var user domain.User
	res := repo.Database.First(&user, "phone = ?", phone)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, domain.ErrUserNotFound
		}
		return nil, res.Error
	}
	return &user, nil
}
