package prices


type Repository interface {
	CreatePrice(Price) error
}

type Service interface {
	CreatePrice(...Price)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) CreatePrice(price ...Price)  {
	for _, p := range price {
		_ = s.repo.CreatePrice(p)
	}
}
