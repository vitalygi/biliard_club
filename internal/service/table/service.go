package table

import (
	"biliard_club/domain"
	"biliard_club/domain/models"
	"errors"
	"fmt"
)

type Service struct {
	Repository domain.TableRepository
}

func NewService(repo domain.TableRepository) *Service {
	return &Service{Repository: repo}
}

func (srv *Service) GetAllTables() ([]models.Table, error) {
	tables, err := srv.Repository.GetAll()
	if err != nil {
		return nil, domain.ErrInternalServer
	}
	return tables, err
}

func (srv *Service) GetByID(id uint) (*models.Table, error) {
	table, err := srv.Repository.GetByID(id)
	if err != nil && errors.Is(err, domain.ErrNotFound) {
		return nil, domain.ErrInternalServer
	}
	return table, err
}

func (srv *Service) Create(table *models.Table) (*models.Table, error) {
	table, err := srv.Repository.Create(table)
	if err != nil && !errors.Is(err, domain.ErrConflict) {
		return nil, domain.ErrInternalServer
	}
	return table, err
}

func (srv *Service) Update(table *models.Table) error {
	err := srv.Repository.Update(table)
	if err != nil && !errors.Is(err, domain.ErrNotFound) {
		fmt.Println("gdfgdsgdfg")
		return domain.ErrInternalServer
	}
	return err
}
