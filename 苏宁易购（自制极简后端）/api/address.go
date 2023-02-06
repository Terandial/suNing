package api

import (
	"github.com/gin-gonic/gin"
	"holidy/awesomeProject/middleware"
	"holidy/awesomeProject/servive"
	"holidy/awesomeProject/tools"
)

func addAddress(c *gin.Context) {
	id, err := middleware.CheckLogin(c)
	if err != nil {
		tools.RespUnauthorized(c)
		return
	}
	place := c.PostForm("address")
	name := c.PostForm("name")
	phone := c.PostForm("phone")
	if place == "" {
		tools.RespFail(c, "收货地址不能为空")
		return
	}
	if name == "" {
		tools.RespFail(c, "收货人不能为空")
		return
	}
	if phone == "" {
		tools.RespFail(c, "手机号不能为空")
		return
	}
	err = servive.CreateAddress(id, place, name, phone)
	if err != nil {
		tools.RespFail(c, "创建地址失败")
		return
	}
	tools.RespSuccess(c, "创建地址成功")
}

func delAddress(c *gin.Context) {
	id, err := middleware.CheckLogin(c)
	if err != nil {
		tools.RespUnauthorized(c)
		return
	}
	aid := c.PostForm("aid")
	if aid == "" {
		tools.RespFail(c, "aid不能为空")
		return
	}
	err = servive.DelAddress(id, aid)
	if err != nil {
		tools.RespFail(c, "删除地址失败")
		return
	}
	tools.RespSuccess(c, "删除地址成功")
}

func viewAddress(c *gin.Context) {
	id, err := middleware.CheckLogin(c)
	if err != nil {
		tools.RespUnauthorized(c)
		return
	}
	address, err := servive.GetAddress(id)
	if err != nil {
		tools.RespFail(c, "未查询到地址")
		return
	}
	c.JSON(200, address)
}
