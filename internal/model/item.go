package model

type Item struct {
	Id               int64
	Order            int64
	Name             string
	Count            int64
	MainShelf        string
	AdditionalShelfs []string
}

type Product_Count struct {
	Product_id int64
	Count      int64
}
