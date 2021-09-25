package models

type CategoryList struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	ParentId uint   `json:"parentId"`
	Level    int    `json:"level"`
	Sort     int    `json:"sort"`
}
