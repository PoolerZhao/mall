package initialize

import (
	"github.com/gin-gonic/gin"
	"mall.com/api"
	"mall.com/global"
	"mall.com/middleware"
)

func Router() {

	engine := gin.Default()

	// 开启跨域、JWT认证
	engine.Use(middleware.Cors())

	// 静态资源请求映射
	engine.Static("/image", global.Config.Upload.SavePath)

	// 用户登录API
	engine.GET("/captcha", api.GetCaptcha)
	engine.POST("/login", api.UserLogin)

	// 开启JWT认证,以下接口需要认证成功才能访问
	engine.Use(middleware.JwtAuth())

	// 文件上传API
	engine.POST("/upload", api.FileUpload)

	// 品牌管理API
	engine.POST("/brand/create", api.CreateBrand)
	engine.DELETE("/brand/delete", api.DeleteBrand)
	engine.PUT("/brand/update", api.UpdateBrand)
	engine.GET("/brand/list", api.GetBrandList)
	engine.GET("/brand/option", api.GetBrandOption)

	// 类目管理API
	engine.POST("/category/create", api.CreateCategory)
	engine.DELETE("/category/delete", api.DeleteCategory)
	engine.PUT("/category/update", api.UpdateCategory)
	engine.GET("/category/list", api.GetCategoryList)
	engine.GET("/category/option", api.GetCategoryOption)

	// 商品管理API
	engine.POST("/product/create", api.CreateProduct)
	engine.DELETE("/product/delete", api.DeleteProduct)
	engine.PUT("/product/update", api.UpdateProduct)
	engine.GET("/product/info", api.GetProductInfo)
	engine.GET("/product/list", api.GetProductList)

	// 订单管理API
	engine.DELETE("/order/delete", api.DeleteOrder)
	engine.PUT("/order/update", api.UpdateOrder)
	engine.GET("/order/list", api.GetOrderList)
	engine.GET("/order/detail", api.GetOrderDetail)

	// 用户管理API
	engine.DELETE("/user/delete", api.DeleteUser)
	engine.PUT("/user/update", api.UpdateUser)
	engine.POST("/user/pwd/change", api.ChangePassword)
	engine.GET("/user/list", api.GetUserList)

	// 启动、监听端口
	post := ":" + global.Config.Server.Post
	_ = engine.Run(post)
}
