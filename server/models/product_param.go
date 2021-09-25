package models

type ProductParam struct {
	Id           uint    `form:"id"            json:"id"`
	CategoryId   uint    `form:"categoryId"    json:"categoryId"`
	Kind         int     `form:"kind"          json:"kind"`
	Title        string  `form:"title"         json:"title"`
	BrandId      uint    `form:"brandId"       json:"brandId"`
	Name         string  `form:"name"          json:"name"`
	Price        float64 `form:"price"         json:"price"`
	Amount       int     `form:"amount"        json:"amount"`
	ImageUrl     string  `form:"imageUrl"      json:"imageUrl"`
	SendAddress  string  `form:"sendAddress"   json:"sendAddress"`
	ParcelType   string  `form:"parcelType"    json:"parcelType"`
	SalesService string  `form:"salesService"  json:"salesService"`
	CreatorId    uint    `form:"creatorId"     json:"creatorId"`
	Status       int     `form:"status"        json:"status"`
}
