package api

import (
	"github.com/gin-gonic/gin"
	"mall.com/global"
	"mall.com/response"
)

// FileUpload 文件上传
func FileUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.Failed("上传图片出错", c)
	}
	image := global.Config.Upload
	err = c.SaveUploadedFile(file, image.SavePath + file.Filename)
	if err != nil {
		return
	}
	imageURL := image.AccessUrl + file.Filename
	response.Success("上传图片成功", imageURL, c)
}

