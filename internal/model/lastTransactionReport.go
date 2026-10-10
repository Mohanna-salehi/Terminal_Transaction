package model

import "time"

type LastTransaction struct {
	TerminalId    uint      `json:"terminal_id"`
	TerminalName  string    `json:"terminal_name"`
	TransactionId uint      `json:"transaction_id"`
	Amount        int64     `json:"amount"`
	Date          time.Time `json:"date"`
}
