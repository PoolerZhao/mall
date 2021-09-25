package models

type OrderList struct {
	Id            uint    `json:"id"`
	PaymentStatus int     `json:"paymentStatus"`
	Status        string  `json:"status"`
	TotalPrice    float64 `json:"totalPrice"`
	Created       string  `json:"created"`
}
