package repository

import (
	"context"

	"Terminal_Transaction/internal/model"

	"gorm.io/gorm"
)

type Terminal = model.Terminal

func AddTerminal(ctx context.Context, db *gorm.DB, t *Terminal) error {
	return gorm.G[Terminal](db).Create(ctx, t)
}

func FindAllTerminal(ctx context.Context, db *gorm.DB) ([]Terminal, error) {

	return gorm.G[Terminal](db).Order("id ASC").Find(ctx)
}
