package model

type Item struct {
	OrderId   int64  `db:"order_id"`
	ProductId int64  `db:"product_id"`
	Name      string `db:"name"`
	Count     int64  `db:"count"`
	MainShelf string `db:"shelf"`
}
