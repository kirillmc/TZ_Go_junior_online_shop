package repository

import (
	"context"
	"github.com/kirillmc/TZ_Go_junior_online_shop/internal/model"
)

type ShopRepository interface {
	//GetProductById(ctx context.Context, id int64) (*model.Item, error)
	//GetProductsByOrder(ctx context.Context, orderId int64) ([]int64, error)

	//GetShelfByOrder(ctx context.Context, orderId int64) ([]string, error)
	GetProductsFromOrders(ctx context.Context, orders string) (*[]model.Item, error)
	GetAddShelfsFromProduct(ctx context.Context, productId int64) ([]string, error)
}

/**
SELECT orders_products.order_id, orders_products.product_id, products.name, orders_products.count, shelfs.name as shelf, shelfs.is_main
FROM orders INNER join orders_products ON orders.id = orders_products.order_id
INNER join products ON orders_products.product_id = products.id
INNER join shelfs ON products.id = shelfs.product_id WHERE orders.id = 10 order by shelfs.name
*/
/**
SELECT orders_products.order_id, orders_products.product_id, products.name, orders_products.count, shelfs.name as shelf, shelfs.is_main
FROM orders INNER join orders_products ON orders.id = orders_products.order_id
INNER join products ON orders_products.product_id = products.id
INNER join shelfs ON products.id = shelfs.product_id WHERE orders.id in (10,11,14,15) and shelfs.is_main=true
*/

// CREATE INDEX shelf_name ON shelfs (name);
