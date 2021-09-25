package service

import (
	"mall.com/common"
	"mall.com/global"
	"mall.com/models"
)

type Category struct {
	Id       uint   `gorm:"primaryKey"`
	Name     string `gorm:"name"`
	ParentId uint   `gorm:"parent_id"`
	Level    int    `gorm:"level"`
	Sort     int    `gorm:"sort"`
	Created  string `gorm:"created"`
	Updated  string `gorm:"updated"`
}

// Create 创建类目
func (c *Category) Create(param models.CategoryParam) uint {
	var category Category
	result := global.Db.Where("name = ?", param.Name).First(&category)
	if result.RowsAffected > 0 {
		return category.Id
	}
	category = Category{
		Name:     param.Name,
		ParentId: param.ParentId,
		Level:    param.Level,
		Sort:     param.Sort,
		Created:  common.NowTime(),
	}
	if global.Db.Create(&category).RowsAffected > 0 {
		return category.Id
	}
	return 0
}

// Delete 删除类目
func (c *Category) Delete(id uint) int64 {
	return global.Db.Delete(&Category{}, id).RowsAffected
}

// Update 更新类目
func (c *Category) Update(param models.CategoryParam) int64 {
	category := Category{
		Id:      param.Id,
		Name:    param.Name,
		Sort:    param.Sort,
		Updated: common.NowTime(),
	}
	return global.Db.Model(&category).Updates(category).RowsAffected
}

// GetList 获取类目列表
func (c *Category) GetList(page models.Page, param models.CategoryParam) ([]models.CategoryList, int64) {
	categoryList := make([]models.CategoryList, 0)
	query := &Category{
		Id:       param.Id,
		Name:     param.Name,
		Level:    param.Level,
		ParentId: param.ParentId,
	}
	rows := common.RestPage(page, "category", query, &categoryList, &[]Category{})
	return categoryList, rows
}

// GetOption 获取类目选项
func (c *Category) GetOption() (option []models.CategoryOption) {
	selectList := make([]models.CategoryList, 0)
	global.Db.Table("category").Find(&selectList)
	return getTreeOptions(0, selectList)
}

// 获取树形结构的选项
func getTreeOptions(id uint, cateList []models.CategoryList) (option []models.CategoryOption) {
	optionList := make([]models.CategoryOption, 0)
	for _ , opt := range cateList {
		if opt.ParentId == id && (opt.Level == 1 || opt.Level == 2 || opt.Level == 3) {
			option := models.CategoryOption{
				Value:    opt.Id,
				Label:    opt.Name,
				Children: getTreeOptions(opt.Id, cateList),
			}
			if opt.Level == 3 { option.Children = nil }
			optionList = append(optionList, option)
		}
	}
	return optionList
}
