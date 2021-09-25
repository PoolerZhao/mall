package models

type OrderDetail struct {

	// 订单信息
	Id            uint   `json:"id"`
	Created       string `json:"created"`
	PaymentStatus int    `json:"paymentStatus"`
	Status        string `json:"status"`

	// 地址信息
	Name            string `json:"name"`
	Mobile          string `json:"mobile"`
	Province        string `json:"province"`
	City            string `json:"city"`
	District        string `json:"district"`
	DetailedAddress string `json:"detailedAddress"`

	// 商品列表信息
	ProductItem []ProductItem `json:"productItem"`
}
