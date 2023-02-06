package model

type User struct {
	Id       int    `json:"id"` //递增排序
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    int    `json:"phone"`
}
