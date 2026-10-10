package model

type Terminal struct {
	Id   uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Code string `json:"code" gorm:"uniqueKey"`
	Name string `json:"name"`
}
