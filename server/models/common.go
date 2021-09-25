package models

// 主键请求参数
type PrimaryKey struct {
	Id uint `form:"id" json:"id"`
}

// 页面请求参数
type Page struct {
	PageNum  int `form:"pageNum"  json:"pageNum"`
	PageSize int `form:"pageSize" json:"pageSize"`
}