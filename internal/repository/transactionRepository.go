package repository

import (
	"context"
	"time"

	"Terminal_Transaction/internal/model"

	"gorm.io/gorm"
)

type Transaction = model.Transaction

func AddTransaction(ctx context.Context, db *gorm.DB, t Transaction) error {
	return gorm.G[Transaction](db).Create(ctx, &t)
}

func FindAllTransaction(ctx context.Context, db *gorm.DB) ([]Transaction, error) {
	return gorm.G[Transaction](db).
		Preload("Terminal", func(db gorm.PreloadBuilder) error {
			return nil
		}).
		Order("terminal_id ASC").
		Find(ctx)
}

func FindTransactionByDate(ctx context.Context, db *gorm.DB, start time.Time, end time.Time) ([]Transaction, error) {

	return gorm.G[Transaction](db).
		Preload("Terminal", func(db gorm.PreloadBuilder) error {
			return nil
		}).
		Where(
			"date >= ? AND date < ?",
			start,
			end,
		).
		Order("terminal_id ASC").
		Find(ctx)
}

func FindTransactionByTerminalId(ctx context.Context, db *gorm.DB, terminalId uint) ([]Transaction, error) {
	return gorm.G[Transaction](db).Where("terminal_id = ?", terminalId).Find(ctx)
}
