package entity

import "time"

type Product struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `json:"title"`
	Price     int    `json:"price"`
	CreatedAt time.Time
	UpdateAt  time.Time
}
