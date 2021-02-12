package entity

type User struct {
	Id       int    `json:"-" db:"id"`
	Username string `form:"username" binding:"required,min=4,max=25,alphanum"`
	Password string `form:"password" binding:"required,min=4,max=25,alphanum"`
}
