package prices

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)


type Repository interface {
	Create(Price) error
	GetAllPrices() ([]*Price, error) 
}

type Service interface {
	CreatePrice(...Price)
	ReadData() error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) CreatePrice(price ...Price)  {
	for _, p := range price {
		_ = s.repo.Create(p)
	}
}

func (s *service) ReadData() error {
	prices, err := s.repo.GetAllPrices()
	if err != nil {
		return err
	}

	log.Info("qtty of prices in DB: ", len(prices))
	log.Info("showing top 100 rows:")

	for i := 0; i < 100; i++ {
		fmt.Printf("ID: %d, ProdNumber: %s\n", prices[i].ID, prices[i].ProductNumber)
	}
  
	return nil
}
