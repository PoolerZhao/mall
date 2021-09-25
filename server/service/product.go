package service

import (
	"mall.com/common"
	"mall.com/global"
	"mall.com/models"
)

type Product struct {
	Id           uint    `gorm:"primaryKey"`
	CategoryId   uint    `gorm:"category_id"`
	Kind         int     `gorm:"kind"`
	Title        string  `gorm:"title"`
	BrandId      uint    `gorm:"brand_id"`
	Name         string  `gorm:"name"`
	Price        float64 `gorm:"price"`
	Amount       int     `gorm:"amount"`
	Sales        int     `gorm:"sales"`
	ImageUrl     string  `gorm:"image_url"`
	SendAddress  string  `gorm:"send_address"`
	ParcelType   string  `gorm:"parcel_type"`
	SalesService string  `gorm:"sales_service"`
	CreatorId    uint    `gorm:"creator_id"`
	Status       int     `gorm:"status"`
	Created      string  `gorm:"created"`
	Updated      string  `gorm:"updated"`
}

// Create 创建商品
func (p *Product) Create(param models.ProductParam) int64 {
	product := Product{
		CategoryId:     param.CategoryId,
		Kind:           param.Kind,
		Title:          param.Title,
		BrandId:        param.BrandId,
		Name:           param.Name,
		Price:          param.Price,
		Amount:         param.Amount,
		ImageUrl:       param.ImageUrl,
		SendAddress:    param.SendAddress,
		ParcelType:     param.ParcelType,
		SalesService:   param.SalesService,
		CreatorId:      param.CreatorId,
		Status:         param.Status,
		Created:        common.NowTime(),
	}
	return global.Db.Create(&product).RowsAffected
}

// Delete 删除商品
func (p *Product) Delete(id uint) int64 {
	return global.Db.Delete(&Product{}, id).RowsAffected
}

// Update 更新商品
func (p *Product) Update(param models.ProductParam) int64 {
	product := Product{
		Id: 			param.Id,
		CategoryId:     param.CategoryId,
		Kind:           param.Kind,
		Title:          param.Title,
		BrandId:        param.BrandId,
		Name:           param.Name,
		Price:          param.Price,
		Amount:         param.Amount,
		ImageUrl:       param.ImageUrl,
		SendAddress:    param.SendAddress,
		ParcelType:     param.ParcelType,
		SalesService:   param.SalesService,
		CreatorId:      param.CreatorId,
		Status:         param.Status,
		Updated:        common.NowTime(),
	}
	return global.Db.Model(&product).Where("id = ?", param.Id).Updates(product).RowsAffected
}

// GetInfo 获取商品信息
func (p *Product) GetInfo(id uint) models.ProductInfo {
	var product models.ProductInfo
	global.Db.Table("product").First(&product, id)
	return product
}

// GetList 获取商品列表
func (p *Product) GetList(page models.Page, param models.ProductParam) ([]models.ProductList, int64) {
	query := &Product{
		Id:         param.Id,
		CategoryId: param.CategoryId,
		Title:      param.Title,
		BrandId:    param.BrandId,
		Status:     param.Status,
	}
	productList := make([]models.ProductList, 0)
	rows := common.RestPage(page, "product", query, &productList, &[]Product{})
	return productList, rows
}

