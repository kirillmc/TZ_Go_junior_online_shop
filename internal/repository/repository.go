package repository

import (
	"context"

	"github.com/kirillmc/TZ_Go_junior_online_shop/internal/model"
)

type ShopRepository interface {
	GetProductsFromOrders(ctx context.Context, orders string) (*[]model.Item, error)
	GetAddShelfsFromProduct(ctx context.Context, productId int64) ([]string, error)
}
