package shop

import "context"

func (i *Implementation) PrintOrderByNumber(ctx context.Context, numbers []int64) error {
	err := i.shopService.PrintOrderByNumber(ctx, numbers)
	if err != nil {
		return err
	}

	return nil
}
