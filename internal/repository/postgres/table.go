package postgres

import (
	"biliard_club/domain"
	"biliard_club/pkg/db"
	"errors"
	"gorm.io/gorm"
)

type TableRepository struct {
	Database *db.Db
}

func NewTableRepository(database *db.Db) *TableRepository {
	return &TableRepository{
		Database: database,
	}
}

func (repo *TableRepository) Create(table *domain.Table) (*domain.Table, error) {
	res := repo.Database.Create(table)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrDuplicatedKey) {
			return nil, domain.ErrTableExists
		}
		return nil, res.Error
	}
	return table, nil
}

func (repo *TableRepository) GetByID(id uint) (*domain.Table, error) {
	var table domain.Table
	res := repo.Database.First(&table, "id = ?", id)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, domain.ErrNotFound
		} else {
			return nil, res.Error
		}
	}
	return &table, nil
}

func (repo *TableRepository) GetAll() ([]domain.Table, error) {
	var tables []domain.Table
	res := repo.Database.Find(&tables)
	if res.Error != nil {
		return tables, res.Error
	}
	return tables, nil
}

func (repo *TableRepository) Update(table *domain.Table) error {
	res := repo.Database.Save(table)
	return res.Error
}
