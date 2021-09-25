package models

type CategoryParam struct {
	Id       uint   `form:"id"       json:"id"`
	Name     string `form:"name"     json:"name"`
	ParentId uint   `form:"parentId" json:"parentId"`
	Level    int    `form:"level"    json:"level"`
	Sort     int    `form:"sort"     json:"sort"`
}
