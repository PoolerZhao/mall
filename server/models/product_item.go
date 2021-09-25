package models

type ProductItem struct {
	Id       uint    `gorm:"primaryKey" json:"id"`
	ImageUrl string  `gorm:"imageUrl"   json:"imageUrl"`
	Name     string  `gorm:"name"       json:"name"`
	Price    float64 `gorm:"price"      json:"price"`
}
