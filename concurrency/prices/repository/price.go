package repository

import "time"

type Price struct {
	ID            int `gorm:"primaryKey"`
	ProductNumber string
	StoreNumber   string
	Price         string
	SalePrice     string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
