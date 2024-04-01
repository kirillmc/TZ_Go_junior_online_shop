package service

import "context"

type ShopService interface {
	PrintOrderByNumber(ctx context.Context, numbers []int64) error
}
