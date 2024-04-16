package models

type Product struct {
	Id               int64            `json:"id" gorm:"primarykey"`
	Name             string           `json:"name"`
	Stock            string           `json:"stock"`
	Price            string           `json:"price"`
	StockDecreaseLog StockDecreaseLog `gorm:"foreginkey:ProductRefer"`
}

