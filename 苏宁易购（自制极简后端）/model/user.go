package model

type User struct {
	Id       int    `json:"id"` //ιε’ζεΊ
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    int    `json:"phone"`
}
