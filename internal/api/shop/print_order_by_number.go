package shop

import "context"

func (i *Implementation) PrintOrderByNumber(ctx context.Context, orders string) error {
	err := i.shopService.PrintOrderByNumber(ctx, orders)
	if err != nil {
		return err
	}

	return nil
}
