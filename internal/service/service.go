package service

import "context"

type ShopService interface {
	PrintOrderByNumber(ctx context.Context, orders string) error
}
