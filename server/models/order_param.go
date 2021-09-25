package models

type OrderParam struct {
	Id            uint    `form:"id"             json:"id"`
	Amount        float64 `form:"amount"         json:"amount"`
	PaymentStatus int     `form:"paymentStatus"  json:"paymentStatus"`
	ProductCount  int     `form:"productCount"   json:"productCount"`
	Total         float64 `form:"total"          json:"total"`
	Status        string  `form:"status"         json:"status"`
	UserId        uint    `form:"userId"         json:"userId"`
	ProductId     uint    `form:"productId"      json:"productId"`
	AddressId     uint    `form:"addressId"      json:"addressId"`
}
