package table

import (
	"biliard_club/pkg/db"
	"errors"
	"gorm.io/gorm"
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

func (repo *Repository) Create(table *Table) (*Table, error) {
	res := repo.Database.Create(table)
	if res.Error != nil {
		slog.Error("failed to create table",
			slog.Group("table",
				"id", table.ID))
	}
	return table, res.Error
}

func (repo *Repository) GetByID(id uint) (*Table, error) {
	var table Table
	res := repo.Database.First(&table, "id = ?", id)
	if res.Error != nil {
		slog.Warn("failed to get table by id",
			slog.Group("table",
				"id", id))
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, NotFound
		} else {
			return nil, res.Error
		}
	}
	return &table, nil
}

func (repo *Repository) GetAll() ([]Table, error) {
	var tables []Table
	res := repo.Database.Find(&tables)
	if res.RowsAffected == 0 {
		return tables, NoTables
	}
	return tables, nil
}

func (repo *Repository) Update(table *Table) error {
	res := repo.Database.Save(table)
	return res.Error
}
