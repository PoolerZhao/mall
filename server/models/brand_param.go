package models

type BrandParam struct {
	Id   uint   `form:"id"    json:"id"`
	Name string `form:"name"  json:"name"`
	Sort int    `form:"sort"  json:"sort"`
}
