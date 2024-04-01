package model

type Item struct {
	Id               int64
	Order            int64
	Name             string
	Count            int64
	MainShelf        string
	AdditionalShelfs []string
}
