package table

import (
	"biliard_club/domain"
)

type TableRepository interface {
	GetAll() ([]domain.Table, error)
	GetByID(id uint) (*domain.Table, error)
	Create(table *domain.Table) (*domain.Table, error)
	Update(table *domain.Table) error
}

type Service struct {
	Repository TableRepository
}

func NewService(repo TableRepository) *Service {
	return &Service{Repository: repo}
}

func (srv *Service) GetAllTables() ([]domain.Table, error) {
	tables, err := srv.Repository.GetAll()
	return tables, err
}

func (srv *Service) GetByID(id uint) (*domain.Table, error) {
	table, err := srv.Repository.GetByID(id)
	return table, err
}

func (srv *Service) Create(table *domain.Table) (*domain.Table, error) {
	return srv.Repository.Create(table)
}

func (srv *Service) Update(table *domain.Table) error {
	return srv.Repository.Update(table)
}
