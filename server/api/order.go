package api

import (
	"github.com/gin-gonic/gin"
	"mall.com/models"
	"mall.com/response"
	"mall.com/service"
)

var order service.Order

// DeleteOrder 删除订单
func DeleteOrder(c *gin.Context) {
	var key models.PrimaryKey
	_ = c.Bind(&key)
	count := order.Delete(key.Id)
	if count > 0 {
		response.Success("删除成功", count, c)
	}
	response.Failed("删除失败", c)
}

// UpdateOrder 更新订单
func UpdateOrder(c *gin.Context) {
	var param models.OrderParam
	_ = c.ShouldBindJSON(param)
	count := order.Update(param)
	if count > 0 {
		response.Success("更新成功", count, c)
	}
	response.Failed("更新失败", c)
}

// GetOrderList 获取订单列表
func GetOrderList(c *gin.Context) {
	var page models.Page
	_ = c.Bind(&page)
	orderList, row := order.GetList(page)
	response.SuccessPage("操作成功", orderList, row, c)
}

// GetOrderDetail 获取订单详情
func GetOrderDetail(c *gin.Context) {
	var param models.OrderParam
	_ = c.Bind(&param)
	orderDetail := order.GetDetail(param.Id)
	response.Success("操作成功", orderDetail, c)
}