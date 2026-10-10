package report

import (
	"Terminal_Transaction/internal/model"
	"Terminal_Transaction/internal/repository"
	"context"
	"time"

	"gorm.io/gorm"
)

func DailyTransactions(ctx context.Context, db *gorm.DB) ([]model.TransactionReport, error) {
	today := time.Now()

	start := time.Date(
		today.Year(),
		today.Month(),
		today.Day(),
		0,
		0,
		0,
		0,
		today.Location(),
	)
	end := start.AddDate(0, 0, 1)

	date, err := repository.FindTransactionByDate(ctx, db, start, end)
	if err != nil {
		return nil, err
	}
	tr := make([]model.TransactionReport, 0, len(date))
	for _, t := range date {
		tr = append(tr, model.TransactionReport{
			Id:           t.Id,
			Amount:       t.Amount,
			Date:         t.Date,
			TerminalName: t.Terminal.Name,
			TerminalId:   t.Terminal.Id,
		})
	}

	return tr, nil
}
