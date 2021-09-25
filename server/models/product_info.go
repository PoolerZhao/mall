package models

type ProductInfo struct {
	Id           uint    `json:"id"`
	CategoryId   uint    `json:"categoryId"`
	Kind         int     `json:"kind"`
	Title        string  `json:"title"`
	BrandId      uint    `json:"brandId"`
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	Amount       int     `json:"amount"`
	Sales        int     `json:"sales"`
	ImageUrl     string  `json:"imageUrl"`
	SendAddress  string  `json:"sendAddress"`
	ParcelType   string  `json:"parcelType"`
	SalesService string  `json:"salesService"`
}
