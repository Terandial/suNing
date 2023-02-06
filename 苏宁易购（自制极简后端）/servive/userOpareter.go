package servive

import (
	"fmt"
	"holidy/awesomeProject/dao"
	"holidy/awesomeProject/model"
	"strconv"
)

// 创建用户
func CreatUser(username string, password string, phone string) error {

	err := dao.InsertUser(username, password, phone)
	return err
}

// 修改密码
func Change(username, newpass string) error {
	err := dao.ChangePassword(username, newpass)
	return err
}

// 判断用户名是否重复
func CheckExist(username string) (flag bool) {
	u, _ := dao.SearchUserByUsername(username)
	flag = true
	if u.Username != username {
		flag = false
		return flag
	}
	return flag
}
func GetUserById(id int) (u model.User) {
	u, err := dao.SearchUserById(id)
	if err != nil {
		fmt.Printf("在获取时未找到用户，重大错误:%v", err)
		return
	}
	return u

}
func CreateWallet(id int) (err error) {
	err = dao.InsertWallet(id)
	return err
}
func GetUser(username string) (u model.User) {
	u, err := dao.SearchUserByUsername(username)
	if err != nil {
		fmt.Printf("在获取时未找到用户，重大错误:%v", err)
		return
	}
	return u
}

func CreateUserInformation(id int, phone string) (err error) {
	intPhone, _ := strconv.Atoi(phone)
	err = dao.InsertUserInformation(id, intPhone)
	return err
}

func CheckUsername(username string) (flag bool) {
	flag = true
	_, err := dao.SearchUserByUsername(username)
	if err != nil {
		flag = false
		return flag
	}
	return flag
}

func VerifyPassword(username string, password string) (flag bool) {
	flag = true
	u, err := dao.SearchUserByUsername(username)
	if err != nil {
		fmt.Printf("在验证密码时未找到用户，重大错误:%v", err)
		return
	}
	if u.Password != password {
		flag = false
		return flag
	}
	return flag
}
func SearchUserByUsername(name string) (u model.User, err error) {
	u, err = dao.SearchUserByUsername(name)
	return u, err
}

func ReviseInformation(id int, nickname string, phone int, gender string, age int) (err error) {
	err = dao.UpdateUserInformation(id, nickname, phone, gender, age)
	return err
}
func ChangePhone(id int, phone int) (err error) {
	err = dao.ChangePhone(id, phone)
	return err
}
