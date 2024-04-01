package shop

import "github.com/kirillmc/TZ_Go_junior_online_shop/internal/service"

type Implementation struct {
	shopService service.ShopService
}

func NewImplementation(shopService service.ShopService) *Implementation {
	return &Implementation{
		shopService: shopService,
	}
}
