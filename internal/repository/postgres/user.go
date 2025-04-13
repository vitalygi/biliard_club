package postgres

import (
	"biliard_club/domain"
	"biliard_club/domain/models"
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

func (repo *UserRepository) Create(user *models.User) (*models.User, error) {
	result := repo.Database.Create(user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, domain.ErrConflict
		}
		return nil, result.Error
	}
	return user, nil
}

func (repo *UserRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	res := repo.Database.First(&user, "id = ?", id)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, domain.ErrNotFound
		}
		return nil, res.Error
	}
	return &user, nil
}

func (repo *UserRepository) GetByPhone(phone string) (*models.User, error) {
	var user models.User
	res := repo.Database.First(&user, "phone = ?", phone)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, domain.ErrNotFound
		}
		return nil, res.Error
	}
	return &user, nil
}

func (repo *UserRepository) Update(user *models.User) error {
	res := repo.Database.Where("id = ?", user.ID).Updates(user)
	if res.Error == nil && res.RowsAffected == 0 {
		return domain.ErrNotFound
	}
	return res.Error
}
