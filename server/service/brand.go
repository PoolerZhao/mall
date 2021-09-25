package service

import (
	"mall.com/common"
	"mall.com/global"
	"mall.com/models"
)

type Brand struct {
	Id      uint   `gorm:"primaryKey"`
	Name    string `gorm:"name"`
	Sort    int    `gorm:"sort"`
	Created string `gorm:"created"`
	Updated string `gorm:"updated"`
}

// Create 创建品牌
func (b *Brand) Create(param models.BrandParam) int64 {
	brand := Brand{
		Name:    param.Name,
		Sort:    param.Sort,
		Created: common.NowTime(),
	}
	return global.Db.Select("Name", "Sort", "Created").Create(&brand).RowsAffected
}

// Delete 删除品牌
func (b *Brand) Delete(id uint) int64 {
	return global.Db.Delete(&Brand{}, id).RowsAffected
}

// Update 修改品牌
func (b *Brand) Update(param models.BrandParam) int64 {
	brand := Brand{
		Id:      param.Id,
		Name:    param.Name,
		Sort:    param.Sort,
		Updated: common.NowTime(),
	}
	return global.Db.Model(&brand).Updates(brand).RowsAffected
}

// GetList 获取品牌列表
func (b *Brand) GetList(page models.Page, param models.BrandParam) ([]models.BrandList, int64) {
	brandList := make([]models.BrandList, 0)
	query := &Brand{ Name: param.Name }
	rows := common.RestPage(page, "brand", query, &brandList, &[]Brand{})
	return brandList, rows
}

// GetOption 获取品牌选项
func (b *Brand) GetOption() []models.BrandOption {
	options := make([]models.BrandOption, 0)
	global.Db.Table("brand").Find(&options)
	return options
}
