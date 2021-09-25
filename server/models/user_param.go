package models

type UserParam struct {

	Id       uint   `form:"primaryKey" json:"id"`
	Username string `form:"username"   json:"username"`
	Password string `form:"password"   json:"password"`
	Avatar   string `form:"avatar"     json:"avatar"`
	Role     string `form:"role"       json:"role"`
	Status   int    `form:"status"     json:"status"`

	OldPassword string `form:"oldPassword"   json:"oldPassword"`
	NewPassword string `form:"newPassword"   json:"newPassword"`

	CaptchaId    string `form:"captchaId"    json:"captchaId"`
	CaptchaValue string `form:"captchaValue" json:"captchaValue"`
}
