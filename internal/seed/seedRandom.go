package seed

import (
	"Terminal_Transaction/internal/model"
	"context"
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

const (
	terminalCount    = 1000
	transactionCount = 1000000
	tranbatchSize    = 5000
	tebatchSize      = 500
	randTerminal     = 10
)

func Seed(ctx context.Context, db *gorm.DB) error {

	terminals := make([]model.Terminal, 0, terminalCount)

	for i := 0; i < terminalCount; i++ {
		terminals = append(terminals, model.Terminal{
			Code: fmt.Sprintf("CODE_%d", rand.Int63n(100000000)),
			Name: fmt.Sprintf("TER_%d", rand.Int63n(1000000000)),
		})
	}

	err := db.WithContext(ctx).CreateInBatches(&terminals, tebatchSize).Error

	if err != nil {
		return err
	}

	terminalIDs := make([]uint, 0, len(terminals))

	for _, terminal := range terminals {
		terminalIDs = append(terminalIDs, terminal.Id)
	}

	now := time.Now()
	thisWeek := 7 * 24 * time.Hour

	transactions := make([]model.Transaction, 0, tranbatchSize)

	for i := 0; i < transactionCount; i++ {

		randomTerminalID := terminalIDs[rand.Intn(len(terminalIDs))]

		randomDate := now.Add(time.Duration(rand.Int63n(int64(thisWeek))))

		transactions = append(transactions, model.Transaction{
			TerminalId: randomTerminalID,
			Date:       randomDate,
			Amount:     int64(rand.Intn(100000)),
		})

		if len(transactions) == tranbatchSize {
			err := db.WithContext(ctx).CreateInBatches(&transactions, tranbatchSize).Error
			if err != nil {
				return err
			}
			transactions = transactions[:0]
		}
	}

	randomTerminals := make([]model.Terminal, 0, randTerminal)

	for i := 0; i < randTerminal; i++ {
		randomTerminals = append(randomTerminals, model.Terminal{
			Code: fmt.Sprintf("CODE_%d", rand.Int63n(100000000)),
			Name: fmt.Sprintf("TER_%d", rand.Int63n(1000000000)),
		})
	}

	err = db.WithContext(ctx).
		CreateInBatches(&randomTerminals, 5).
		Error

	if err != nil {
		return err
	}
	return nil
}
