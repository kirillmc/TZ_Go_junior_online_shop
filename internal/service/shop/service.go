package shop

import (
	"github.com/kirillmc/TZ_Go_junior_online_shop/internal/repository"
	"github.com/kirillmc/TZ_Go_junior_online_shop/internal/service"
)

type serv struct {
	shopRepository repository.ShopRepository
}

func NewService(shopRepository repository.ShopRepository) service.ShopService {
	return &serv{
		shopRepository: shopRepository,
	}
}
