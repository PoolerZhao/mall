package models

type ProductList struct {
	Id       uint    `json:"id"`
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	Amount   int     `json:"amount"`
	ImageUrl string  `json:"imageUrl"`
	Sales    int     `json:"sales"`
	Status   int     `json:"status"`
	Created  string  `json:"created"`
}
