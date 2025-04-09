package postgres

import (
	"biliard_club/domain"
	"biliard_club/pkg/db"
)

type GameRepository struct {
	Database *db.Db
}

func NewGameRepository(database *db.Db) *GameRepository {
	return &GameRepository{
		Database: database,
	}
}

func (repo *GameRepository) Create(game *domain.Game) error {
	res := repo.Database.Create(game)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (repo *GameRepository) GetById(id uint) (*domain.Game, error) {
	var game domain.Game
	res := repo.Database.
		Preload("Table").
		Preload("User").
		Find(&game, "id = ?", id)
	if res.Error != nil {
		return nil, res.Error
	}
	return &game, nil
}
