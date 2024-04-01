package shop

import (
	"context"
	"fmt"
	"github.com/kirillmc/TZ_Go_junior_online_shop/internal/model"
)

func (s *serv) PrintOrderByNumber(ctx context.Context, orders string) error {
	items := []model.Item{}

	itms, err := s.shopRepository.GetProductsFromOrders(ctx, orders)
	if err != nil {
		return err
	}

	items = append(items, *itms...)

	for i, item := range items {
		addShelfs, err := s.shopRepository.GetAddShelfsFromProduct(ctx, item.Id)
		if err != nil {
			return err
		}

		items[i].AdditionalShelfs = addShelfs
	}

	printer(items)

	return nil
}

func printer(items []model.Item) {
	if len(items) >= 1 {
		s := fmt.Sprintf("===Стеллаж %s\n%s", items[0].MainShelf, strElem(items[0]))
		fmt.Print(s)
	}
	if len(items) > 1 {
		for i := 1; i < len(items); i++ {
			if items[i-1].MainShelf != items[i].MainShelf {
				s := fmt.Sprintf("===Стеллаж %s\n%s", items[i].MainShelf, strElem(items[i]))
				fmt.Print(s)
			} else {
				fmt.Print(strElem(items[i]))
			}
		}
	}
}

func strElem(item model.Item) string {
	s := fmt.Sprintf("%s (id=%d)\nзаказ %d, %d шт\n", item.Name, item.Id, item.Order, item.Count)
	if len(item.AdditionalShelfs) > 0 {
		s = s + fmt.Sprintf("доп стеллаж: ")
		for _, elem := range item.AdditionalShelfs {
			s = s + elem + ","
		}
		s = s[:len(s)-1] + "\n\n"
	} else {
		s = s + "\n"
	}
	return s
}
