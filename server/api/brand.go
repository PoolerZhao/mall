package api

import (
	"github.com/gin-gonic/gin"
	"mall.com/models"
	"mall.com/response"
	"mall.com/service"
)

var brand service.Brand

// CreateBrand 创建品牌
func CreateBrand(c *gin.Context) {
	var param models.BrandParam
	_ = c.ShouldBindJSON(&param)
	count := brand.Create(param)
	if count > 0 {
		response.Success("创建成功", count, c)
		return
	}
	response.Failed("创建失败", c)
}

// DeleteBrand 删除品牌
func DeleteBrand(c *gin.Context) {
	var key models.PrimaryKey
	_ = c.Bind(&key)
	count := brand.Delete(key.Id)
	if count > 0 {
		response.Success("删除成功", count, c)
		return
	}
	response.Failed("删除失败", c)
}

// UpdateBrand 更新品牌
func UpdateBrand(c *gin.Context) {
	var param models.BrandParam
	_ = c.ShouldBindJSON(&param)
	count := brand.Update(param)
	if count > 0 {
		response.Success("更新成功", count, c)
		return
	}
	response.Failed("更新失败", c)
}

// GetBrandList 获取品牌列表
func GetBrandList(c *gin.Context) {
	var page models.Page
	var param models.BrandParam
	_ = c.Bind(&page)
	_ = c.Bind(&param)
	brandList, rows := brand.GetList(page, param)
	response.SuccessPage("操作成功", brandList, rows, c)
}

// GetBrandOption 获取品牌选项
func GetBrandOption(c *gin.Context) {
	options := brand.GetOption()
	response.Success("操作成功", options, c)
}