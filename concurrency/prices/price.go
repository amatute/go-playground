package prices

import "time"

type Price struct {
	ID            int       `json:"id"`
	ProductNumber string    `json:"product_number"`
	StoreNumber   string    `json:"store_number"`
	Price         string    `json:"proce"`
	SalePrice     string    `json:"sale_price"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
