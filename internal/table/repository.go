package table

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

func (repo *Repository) Create(table *Table) error {
	res := repo.Database.Create(table)
	if res.Error != nil {
		slog.Error("failed to create table",
			slog.Group("table",
				"id", table.ID))
	}
	return res.Error
}

func (repo *Repository) GetByID(id uint) (*Table, error) {
	var table Table
	res := repo.Database.Preload("User").First(&table, "id = ?", id)
	if res.Error != nil {
		slog.Error("failed to get table by id",
			slog.Group("table",
				"id", id))
		return nil, res.Error
	}
	return &table, nil
}
