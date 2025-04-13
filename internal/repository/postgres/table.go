package postgres

import (
	"biliard_club/domain"
	"biliard_club/domain/models"
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

func (repo *TableRepository) Create(table *models.Table) (*models.Table, error) {
	res := repo.Database.Create(table)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrDuplicatedKey) {
			return nil, domain.ErrConflict
		}
		return nil, res.Error
	}
	return table, nil
}

func (repo *TableRepository) GetByID(id uint) (*models.Table, error) {
	var table models.Table
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

func (repo *TableRepository) GetAll() ([]models.Table, error) {
	var tables []models.Table
	res := repo.Database.Find(&tables)
	if res.Error != nil {
		return tables, res.Error
	}
	return tables, nil
}

func (repo *TableRepository) Update(table *models.Table) error {
	res := repo.Database.Where("id = ?", table.ID).Updates(table)
	if res.Error == nil && res.RowsAffected == 0 {
		return domain.ErrNotFound
	}
	return res.Error
}
