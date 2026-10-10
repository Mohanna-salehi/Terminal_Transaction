package report

import (
	"Terminal_Transaction/internal/model"
	"Terminal_Transaction/internal/repository"
	"context"

	"gorm.io/gorm"
)

type terminal = model.Terminal

func TerminalReport(ctx context.Context, db *gorm.DB) ([]terminal, error) {
	te, err := repository.FindAllTerminal(ctx, db)
	if err != nil {
		return nil, err
	}

	report := make([]terminal, 0, len(te))
	for _, t := range te {
		report = append(report, terminal{
			Id:   t.Id,
			Code: t.Code,
			Name: t.Name,
		})
	}

	return report, nil
}
