package domain

import "biliard_club/domain/models"

type UserRepository interface {
	GetByPhone(phone string) (*models.User, error)
	GetByID(id uint) (*models.User, error)
	Create(user *models.User) (*models.User, error)
	Update(user *models.User) error
}

type TableRepository interface {
	Create(table *models.Table) (*models.Table, error)
	GetByID(id uint) (*models.Table, error)
	GetAll() ([]models.Table, error)
	Update(table *models.Table) error
}

type GameRepository interface {
	Create(game *models.Game) (*models.Game, error)
	GetByID(id uint) (*models.Game, error)
	Update(user *models.Game) error
}
