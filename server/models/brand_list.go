package models

type BrandList struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Sort    int    `json:"sort"`
	Created string `json:"created"`
}
