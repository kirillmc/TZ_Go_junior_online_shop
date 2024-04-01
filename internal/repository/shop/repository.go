package shop

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kirillmc/TZ_Go_junior_online_shop/internal/model"
	"github.com/kirillmc/TZ_Go_junior_online_shop/internal/repository"
	"github.com/kirillmc/TZ_Go_junior_online_shop/internal/repository/shop/converter"
	modelRepo "github.com/kirillmc/TZ_Go_junior_online_shop/internal/repository/shop/model"
)

const (
	PRODUCT_ID            = "product_id"
	ORDER_ID              = "order_id"
	TABLE_ORDERS_PRODUCTS = "orders_products"
	TABLE_ORDERS          = "orders"
	TABLE_PRODUCTS        = "products"
	TABLE_SHELFS          = "shelfs"

	SHELFS_NAME              = "shefs.name"
	ORDERS_ID                = "orders.id"
	ORDERS_PRODUCTS_ORDER_ID = "orders_products.order_id"
	PRODUCTS_ID              = "products.id"
	SHELFS_PRODUCT_ID        = "shelfs.product_id "
	SHELFS_IS_MAIN           = "shelfs.is_main"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.ShopRepository {
	return &repo{db: db}
}

//	func (r *repo) GetProductById(ctx context.Context, num int64) (*model.Item, error) {
//		var products []modelRepo.Item
//
//		query := fmt.Sprintf("SELECT %s FROM %s WHERE %s=$1", PRODUCT_ID, TABLE_ORDERS_PRODUCTS, ORDER_ID)
//		rows, err := r.db.Query(ctx, query, orderId)
//		if err != nil {
//			return nil, err
//		}
//
//		defer rows.Close()
//
//		for rows.Next() {
//			var id int64
//
//			err = rows.Scan(&id)
//			if err != nil {
//				return nil, err
//			}
//
//			ids = append(ids, id)
//		}
//		return nil, nil
//	}
//
//	func (r *repo) GetProductsByOrder(ctx context.Context, orderId int64) ([]int64, error) {
//		var ids []int64
//
//		query := fmt.Sprintf("SELECT %s FROM %s WHERE %s=$1", PRODUCT_ID, TABLE_ORDERS_PRODUCTS, ORDER_ID)
//		rows, err := r.db.Query(ctx, query, orderId)
//		if err != nil {
//			return nil, err
//		}
//
//		defer rows.Close()
//
//		for rows.Next() {
//			var id int64
//
//			err = rows.Scan(&id)
//			if err != nil {
//				return nil, err
//			}
//
//			ids = append(ids, id)
//		}
//
//		return ids, nil
//	}
//func (r *repo) GetShelfByOrder(ctx context.Context, orderId int64) ([]string, error) {
//	//query := fmt.Sprintf("SELECT %s FROM %s "+
//	//	"INNER join %s ON %s = %s "+
//	//	"INNER join %s ON %s = %s "+
//	//	"INNER join %s ON %s = %s "+
//	//	"WHERE %s = $1 and %s=$2", SHELFS_NAME, TABLE_ORDERS, TABLE_ORDERS_PRODUCTS,
//	//	ORDER_ID, ORDERS_PRODUCTS_ORDER_ID, TABLE_PRODUCTS, ORDERS_PRODUCTS_ORDER_ID, PRODUCTS_ID, TABLE_SHELFS, PRODUCTS_ID, SHELFS_PRODUCT_ID, ORDERS_ID, SHELFS_IS_MAIN)
//
//	query := ""
//
//	rows, err := r.db.Query(ctx, query, orderId, true)
//	if err != nil {
//		return nil, err
//	}
//
//	var shelfs []string
//	for rows.Next() {
//		var shelf string
//
//		err = rows.Scan(&shelf)
//		if err != nil {
//			return nil, err
//		}
//
//		shelfs = append(shelfs, shelf)
//	}
//
//	defer rows.Close()
//	return nil, nil
//}

func (r *repo) GetProductsFromOrder(ctx context.Context, orderId int64) (*[]model.Item, error) {
	query := fmt.Sprintf("SELECT orders_products.order_id, orders_products.product_id, products.name, orders_products.count, shelfs.name "+
		"FROM orders INNER join orders_products ON orders.id = orders_products.order_id "+
		"INNER join products ON orders_products.product_id = products.id "+
		"INNER join shelfs ON products.id = shelfs.product_id WHERE orders.id=$1 AND shelfs.is_main=%t", true)

	rows, err := r.db.Query(ctx, query, orderId)
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

/**
SELECT orders_products.order_id, orders_products.product_id, products.name, orders_products.count, shelfs.name as shelf, shelfs.is_main
FROM orders INNER join orders_products ON orders.id = orders_products.order_id
INNER join products ON orders_products.product_id = products.id
INNER join shelfs ON products.id = shelfs.product_id WHERE shelfs.name = 'А' and shelfs.is_main=true
*/
