package report

import (
	"Terminal_Transaction/internal/model"
	"Terminal_Transaction/internal/repository"
	"context"

	"gorm.io/gorm"
)

func TransReport(ctx context.Context, db *gorm.DB) ([]model.TransactionReport, error) {
	tr, err := repository.FindAllTransaction(ctx, db)
	if err != nil {
		return nil, err
	}

	report := make([]model.TransactionReport, 0, len(tr))

	for _, t := range tr {
		report = append(report, model.TransactionReport{
			Id:           t.Id,
			Amount:       t.Amount,
			Date:         t.Date,
			TerminalName: t.Terminal.Name,
			TerminalId:   t.TerminalId,
		})
	}

	return report, nil
}
