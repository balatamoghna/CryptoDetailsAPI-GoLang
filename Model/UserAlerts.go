package model

import (
	"time"

	"gorm.io/gorm"
)

//Alerts details model
type Alerts struct {
	ID        int     `gorm:"primaryKey;auto_increment" json:"id"`
	Email     string  `json:"email"`
	Currency  string  `json:"currency"`
	Target    float64 `json:"target"`
	Triggered string  `json:"triggered"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
