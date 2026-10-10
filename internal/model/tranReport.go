package model

import "time"

type TransactionReport struct {
	Id           uint      `json:"id"`
	TerminalId   uint      `json:"terminal_id"`
	TerminalName string    `json:"terminal_name"`
	Amount       int64     `json:"amount"`
	Date         time.Time `json:"date"`
}
