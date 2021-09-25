package api

import (
	"github.com/gin-gonic/gin"
	"mall.com/common"
	"mall.com/models"
	"mall.com/response"
	"mall.com/service"
)

var user service.User

func GetCaptcha(c *gin.Context) {
	id, b64s, _ := common.GenerateCaptcha()
	data := map[string]interface{}{ "captchaId": id, "captchaImg": b64s}
	response.Success("操作成功", data, c)
}

func UserLogin(c *gin.Context) {
	var param models.UserParam
	_ = c.ShouldBindJSON(&param)

	// 检查验证码
	if !common.VerifyCaptcha(param.CaptchaId, param.CaptchaValue) {
		response.Failed("验证码错误", c)
		return
	}
	// 生成token
	if user.Login(param) > 0 {
		token, _ := common.GenerateToke(param.Username)
		res := map[string]interface{}{ "token": token}
		response.Success("登录成功", res, c)
		return
	}
	response.Failed("用户名或密码错误", c)
}

func UpdateUser(c *gin.Context) {
	var param models.UserParam
	_ = c.ShouldBindJSON(&param)
	count := user.Update(param)
	if count > 0 {
		response.Success("更新成功", count, c)
	}
	response.Failed("更新失败", c)
}

func ChangePassword(c *gin.Context) {
	var param models.UserParam
	_ = c.ShouldBindJSON(&param)
	count := user.ChangePwd(param)
	if count > 0{
		response.Success("密码更新成功", count, c)
	}
	response.Failed("密码更新失败", c)
}

func DeleteUser(c *gin.Context) {
	var key models.PrimaryKey
	_ = c.Bind(&key)
	count := user.Delete(key.Id)
	if count > 0 {
		response.Success("删除成功", count, c)
	}
	response.Failed("删除失败", c)
}

func GetUserList(c *gin.Context) {
	var page models.Page
	_ = c.Bind(&page)
	userList := user.GetList(page)
	response.Success("操作成功", userList, c)
}

