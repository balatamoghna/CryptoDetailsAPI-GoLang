package model

//Currencies details model
type Currencies struct {
	ID           int     `gorm:"auto_increment" json:"id"`
	Symbol       string  `gorm:"primaryKey" json:"symbol"`
	Name         string  `json:"name"`
	CurrentPrice float64 `json:"current_price"`
	Updated      int64   `gorm:"autoUpdateTime:milli"`
}
