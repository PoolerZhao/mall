package models

type UserList struct {
	Id       uint   `json:"id"`
	Avatar   string `json:"avatar"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Status   int    `json:"status"`
	Created  string `json:"created"`
}
