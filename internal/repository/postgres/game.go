package postgres

import (
	"biliard_club/domain"
	"biliard_club/domain/models"
	"biliard_club/pkg/db"
	"errors"
	"gorm.io/gorm"
)

type GameRepository struct {
	Database *db.Db
}

func NewGameRepository(database *db.Db) *GameRepository {
	return &GameRepository{
		Database: database,
	}
}

func (repo *GameRepository) Create(game *models.Game) error {
	res := repo.Database.Create(game)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrDuplicatedKey) {
			return domain.ErrConflict
		}
		return res.Error
	}
	return nil
}

func (repo *GameRepository) GetById(id uint) (*models.Game, error) {
	var game models.Game
	res := repo.Database.
		Preload("Table").
		Preload("User").
		First(&game, "id = ?", id)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, domain.ErrNotFound
		}
		return nil, res.Error
	}
	return &game, nil
}

func (repo *GameRepository) Update(game *models.Game) error {
	res := repo.Database.Where("id = ?", game.ID).Updates(game)
	if res.Error == nil && res.RowsAffected == 0 {
		return domain.ErrNotFound
	}
	return res.Error
}
