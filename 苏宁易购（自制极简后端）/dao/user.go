package dao

import (
	"fmt"
	"holidy/awesomeProject/model"
)

func InsertUser(username string, password string, phone string) (err error) {
	sqlStr := "insert into user(username,password,phone) values (?,?,?)"
	_, err = DB.Exec(sqlStr, username, password, phone)
	if err != nil {
		fmt.Printf("用户插入失败，err:%v", err)
		return err
	}
	return err
}

func SearchUserById(id int) (u model.User, err error) {
	sqlStr := "select id,username,phone,password from user where id=?"
	err = DB.Get(&u, sqlStr, id)
	if err != nil {
		fmt.Printf("用户查询失败，err:%v", err)
		fmt.Printf("id:%v,username:%v,phone:%v,password:%v", u.Id, u.Username, u.Phone, u.Password)
		return u, err
	}
	return u, err
}
func ChangePassword(username string, newPass string) (err error) {
	sqlStr := "update user set password=? where username=?"
	_, err = DB.Exec(sqlStr, newPass, username)
	if err != nil {
		fmt.Printf("修改密码失败，err：%v", err)
	}
	return err
}
func ChangePhone(id int, phone int) (err error) {
	sqlStr := "update user set phone=? where id=?"
	_, err = DB.Exec(sqlStr, phone, id)
	if err != nil {
		fmt.Printf("修改手机号失败，err：%v", err)
	}
	return err
}
func SearchUserByUsername(username string) (u model.User, err error) {
	sqlStr := "select id,username,phone,password from user where username=?"
	err = DB.Get(&u, sqlStr, username)
	if err != nil {
		fmt.Printf("用户查询失败，err:%v", err)
		fmt.Printf("id:%v,username:%v,phone:%v,password:%v", u.Id, u.Username, u.Phone, u.Password)
		return u, err
	}
	return u, err
}
func InsertUserInformation(id int, phone int) (err error) {
	sqlStr := "insert into userinformation (id,nickname,phone,gender,age) values (?,?,?,?,?)"
	_, err = DB.Exec(sqlStr, id, phone, "未填写", "未填写", 0)
	return err
}

func UpdateUserInformation(id int, nickname string, phone int, gender string, age int) (err error) {
	sqlStr := "update userinformation set nickname=?,phone=?,gender=?,age=? where id=?"
	_, err = DB.Exec(sqlStr, nickname, phone, gender, age, id)
	return err
}

func SelectUserInformation(id int) (u model.UserInformation, err error) {
	sqlStr := "select id,nickname,phone,gender,age from userinformation where id=?"
	err = DB.Get(&u, sqlStr, id)
	if err != nil {
		return model.UserInformation{}, err
	}
	return u, err
}
