package entity

type User struct {
	Id       int    `json:"id"`
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
