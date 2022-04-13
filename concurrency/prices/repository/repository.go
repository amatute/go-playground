package repository

import (
	"fmt"
	"time"

	"github.com/amatute/go-playground/concurrency/prices"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB :db}
}

func (r *Repository) Create(p prices.Price) error  {
	price := &Price{
		ID:            int(uuid.New().ID()),
		ProductNumber: fmt.Sprintf("PROD_%s", p.ProductNumber),
		StoreNumber:   fmt.Sprintf("STORE_%s", p.StoreNumber),
		Price:         "500",
		SalePrice:     "250",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	result := r.DB.Create(price)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *Repository) GetAllPrices() ([]*prices.Price, error) {
	var prices = make([]*prices.Price,0)
	result := r.DB.Find(&prices)
	if result.Error != nil {
		return prices, result.Error
	}
	return prices, nil	
}


