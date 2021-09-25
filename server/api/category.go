package api

import (
	"github.com/gin-gonic/gin"
	"mall.com/models"
	"mall.com/response"
	"mall.com/service"
)

var category service.Category

// CreateCategory 创建类目
func CreateCategory(c *gin.Context) {
	var param models.CategoryParam
	_ = c.ShouldBindJSON(&param)
	count := category.Create(param)
	if count > 0 {
		response.Success("创建成功", count, c)
		return
	}
	response.Failed("创建失败", c)
}

// DeleteCategory 删除类目
func DeleteCategory(c *gin.Context) {
	var key models.PrimaryKey
	_ = c.Bind(&key)
	count := category.Delete(key.Id)
	if count > 0 {
		response.Success("删除成功", count, c)
		return
	}
	response.Failed("删除失败", c)
}

// UpdateCategory 更新类目
func UpdateCategory(c *gin.Context) {
	var param models.CategoryParam
	_ = c.ShouldBindJSON(&param)
	count := category.Update(param)
	if count > 0 {
		response.Success("更新成功", count, c)
		return
	}
	response.Failed("更新失败", c)
}

// GetCategoryList 获取类目列表
func GetCategoryList(c *gin.Context) {
	var page models.Page
	var param models.CategoryParam
	_ = c.Bind(&page)
	_ = c.Bind(&param)
	cateList, rows := category.GetList(page, param)
	response.SuccessPage("操作成功", cateList, rows, c)
}

// GetCategoryOption 获取类目选项
func GetCategoryOption(c *gin.Context) {
	option := category.GetOption()
	response.Success("操作成功", option, c)
}
