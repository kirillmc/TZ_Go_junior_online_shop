package converter

import (
	"github.com/kirillmc/TZ_Go_junior_online_shop/internal/model"
	modelRepo "github.com/kirillmc/TZ_Go_junior_online_shop/internal/repository/shop/model"
)

func ToItemFromRepo(items []modelRepo.Item) *[]model.Item {
	var modelItems []model.Item

	for _, item := range items {
		var modelItem model.Item

		modelItem.Id = item.ProductId
		modelItem.Order = item.OrderId
		modelItem.Name = item.Name
		modelItem.Count = item.Count
		modelItem.MainShelf = item.MainShelf
		modelItem.AdditionalShelfs = nil

		modelItems = append(modelItems, modelItem)
	}
	return &modelItems
}
