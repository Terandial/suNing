package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"holidy/awesomeProject/middleware"
	"holidy/awesomeProject/servive"
	"holidy/awesomeProject/tools"
	"log"
	"strconv"
)

// 用户注册
func register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	phone := c.PostForm("phone")
	if username == "" || password == "" {
		tools.RespFail(c, "用户名或密码为空")
		return
	}
	if phone == "" {
		tools.RespFail(c, "手机号为空")
		return
	}
	if len(phone) != 11 {
		tools.RespFail(c, "手机号格式错误")
		return
	}
	if servive.CheckExist(username) {
		tools.RespFail(c, "该用户名已存在")
		return
	}
	err := servive.CreatUser(username, password, phone)
	if err != nil {
		tools.RespFail(c, "创建用户失败")
		return
	}

	id := servive.GetUser(username).Id

	err = servive.CreateWallet(id)
	if err != nil {
		tools.RespFail(c, "用户创建钱包失败")
		return
	}
	err = servive.CreateUserInformation(id, phone)
	if err != nil {
		tools.RespFail(c, "用户创建账户信息失败")
		return
	}
	tools.RespSuccess(c, "注册成功")
}

// 用户登录
func login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		tools.RespFail(c, "用户名或密码为空")
		return
	}
	if !servive.CheckUsername(username) {
		tools.RespFail(c, "用户名不存在")
		return
	}
	if !servive.VerifyPassword(username, password) {
		tools.RespFail(c, "密码错误")
	}
	uid := servive.GetUser(username).Id
	uidString := strconv.Itoa(uid)
	c.SetCookie("user", uidString, 3600, "/", "", false, true)
	tools.RespSuccess(c, "登录成功")

}

// 修改密码
func ChangePassword(c *gin.Context) {
	username := c.PostForm("name")
	password := c.PostForm("password")
	newpass := c.PostForm("newpassword")
	if username == "" || password == "" {
		tools.RespFail(c, "不能为空")
		return
	}
	_, err := servive.SearchUserByUsername(username)
	if err != nil {
		if err != sql.ErrNoRows {
			tools.RespFail(c, "用户不存在")
		} else {
			log.Printf("search user err: %#v", err)
			tools.RespUnauthorized(c)
			return
		}
		return
	}
	err = servive.Change(username, newpass)
	if err != nil {
		tools.RespFail(c, "修改密码失败！")
		return
	}
	tools.RespSuccess(c, "修改密码成功！")
}
func logout(c *gin.Context) {
	_, err := middleware.CheckLogin(c)
	if err != nil {
		tools.RespUnauthorized(c)
		return
	}
	c.SetCookie("user", "", -1, "/", "", false, true)
	tools.RespSuccess(c, "注销成功")
}

// 传说中的设置
func setting(c *gin.Context) {
	id, err := middleware.CheckLogin(c)
	if err != nil {
		tools.RespUnauthorized(c)
		return
	}
	nickname := c.PostForm("nickname")
	phone := c.PostForm("phone")
	gender := c.PostForm("gender")
	age := c.PostForm("age")
	intAge, _ := strconv.Atoi(age)
	intPhone, _ := strconv.Atoi(phone)
	if gender == "" {
		gender = "未填写"
	}
	if age == "" {
		intAge = 0
	}
	err = servive.ReviseInformation(id, nickname, intPhone, gender, intAge)
	if err != nil {
		tools.RespFail(c, "修改个人信息失败")
		return
	}
	err = servive.ChangePhone(id, intPhone)
	if err != nil {
		tools.RespFail(c, "修改用户信息失败")
		return
	}
	tools.RespSuccess(c, "修改成功")
}
