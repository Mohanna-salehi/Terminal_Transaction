package model

type DailyTerminalReport struct {
	TerminalId       uint   `json:"terminal_id"`
	TerminalCode     string `json:"terminal_code"`
	TerminalName     string `json:"terminal_name"`
	TransactionCount int64  `json:"transaction_count"`
	TotalAmount      int64  `json:"total_amount"`
	TodayTransaction string `json:"today_transaction"`
}
