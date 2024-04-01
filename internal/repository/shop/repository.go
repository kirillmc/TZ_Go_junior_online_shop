package shop

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kirillmc/TZ_Go_junior_online_shop/internal/model"
	"github.com/kirillmc/TZ_Go_junior_online_shop/internal/repository"
	"github.com/kirillmc/TZ_Go_junior_online_shop/internal/repository/shop/converter"
	modelRepo "github.com/kirillmc/TZ_Go_junior_online_shop/internal/repository/shop/model"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.ShopRepository {
	return &repo{db: db}
}

func (r *repo) GetProductsFromOrders(ctx context.Context, orders string) (*[]model.Item, error) {
	query := fmt.Sprintf("SELECT orders_products.order_id, orders_products.product_id, products.name, orders_products.count, shelfs.name "+
		"FROM orders INNER join orders_products ON orders.id = orders_products.order_id "+
		"INNER join products ON orders_products.product_id = products.id "+
		"INNER join shelfs ON products.id = shelfs.product_id WHERE orders.id in (%s) and shelfs.is_main=true ORDER BY shelfs.name", orders)
	log.Printf("orders is %s", orders)
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var items []modelRepo.Item

	for rows.Next() {
		var item modelRepo.Item

		err = rows.Scan(&item.OrderId, &item.ProductId, &item.Name, &item.Count, &item.MainShelf)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}
	return converter.ToItemFromRepo(items), nil
}

func (r *repo) GetAddShelfsFromProduct(ctx context.Context, productId int64) ([]string, error) {
	query := "SELECT shelfs.name FROM shelfs WHERE product_id=$1 and shelfs.is_main=false"

	rows, err := r.db.Query(ctx, query, productId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var addShelfs []string

	for rows.Next() {
		var sh string

		err = rows.Scan(&sh)
		if err != nil {
			return nil, err
		}

		addShelfs = append(addShelfs, sh)
	}

	if len(addShelfs) == 0 {
		return []string{}, nil
	}

	return addShelfs, nil
}
