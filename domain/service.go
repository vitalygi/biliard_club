package domain

import "biliard_club/domain/models"

type UserService interface {
	Create(user *models.User) (*models.User, error)
}

type TableService interface {
	Update(table *models.Table) error
	Create(table *models.Table) (*models.Table, error)
	GetByID(id uint) (*models.Table, error)
	GetAllTables() ([]models.Table, error)
}

type AuthService interface {
	Register(phone, password, name string) (*models.User, error)
	Login(phone, password string) (*models.User, error)
}
