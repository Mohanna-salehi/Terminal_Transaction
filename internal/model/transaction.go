package model

import "time"

type Transaction struct {
	Id         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Amount     int64     `json:"amount"`
	Date       time.Time `json:"date"`
	TerminalId uint      `json:"terminalId"`
	Terminal   Terminal  `json:"terminal" gorm:"foreignKey:TerminalId;references:Id"`
}
