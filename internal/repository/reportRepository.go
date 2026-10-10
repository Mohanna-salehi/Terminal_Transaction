package repository

import (
	"Terminal_Transaction/internal/model"
	"context"

	"gorm.io/gorm"
)

func FindDailyTerminalReport(
	ctx context.Context,
	db *gorm.DB,
) ([]model.DailyTerminalReport, error) {

	var result []model.DailyTerminalReport

	err := db.
		WithContext(ctx).
		Raw("CALL daily_terminal()").
		Scan(&result).
		Error

	if err != nil {
		return nil, err
	}

	return result, nil
}
func FindLastTransactionReport(
	ctx context.Context,
	db *gorm.DB,
) ([]model.LastTransaction, error) {

	var result []model.LastTransaction

	err := db.
		WithContext(ctx).
		Raw("CALL Last_Transaction()").
		Scan(&result).
		Error

	if err != nil {
		return nil, err
	}

	return result, nil
}
