package game

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

func (repo *Repository) Create(game *Game) error {
	res := repo.Database.Create(game)
	if res.Error != nil {
		slog.Error("failed to create game",
			slog.Group("game",
				"id", game.ID,
				"startTime", game.StartTime))
		return res.Error
	}
	return nil
}

func (repo *Repository) GetById(id uint) (*Game, error) {
	var game Game
	res := repo.Database.
		Preload("Table").
		Preload("User").
		Find(&game, "id = ?", id)
	if res.Error != nil {
		slog.Error("failed to get game by id",
			slog.Group("game",
				"id", id))
		return nil, res.Error
	}
	return &game, nil
}
