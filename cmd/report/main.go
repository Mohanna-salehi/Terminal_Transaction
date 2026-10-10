package main

import (
	"Terminal_Transaction/internal/database"
	"Terminal_Transaction/internal/report"
	"context"
	"fmt"
	"log"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	terminalReport, err := report.TerminalReport(ctx, db)
	if err != nil {
		return
	}
	for _, r := range terminalReport {
		fmt.Println(r.Id, r.Code, r.Name)
	}

	reports, err := report.TransReport(ctx, db)
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range reports {
		fmt.Println(
			r.Id,
			r.TerminalId,
			r.TerminalName,
			r.Amount,
		)
	}

	transactions, err := report.DailyTransactions(ctx, db)
	if err != nil {
		log.Fatal(err)
	}
	for _, t := range transactions {
		fmt.Println(
			t.Id,
			t.TerminalId,
			t.TerminalName,
			t.Amount,
			t.Date)
	}

}
