package api

import (
	"github.com/gin-gonic/gin"
	"holidy/awesomeProject/middleware"
	"holidy/awesomeProject/servive"
	"holidy/awesomeProject/tools"
	"strconv"
)

func browse(c *gin.Context) {
	id, err := middleware.CheckLogin(c)
	if err != nil {
		tools.RespUnauthorized(c)
		return
	}
	o, err := servive.SearchOrder(id)
	if err != nil {
		tools.RespFail(c, "未查询到订单")
		return
	}
	c.JSON(200, o)
}
func cancelOrder(c *gin.Context) {
	id, err := middleware.CheckLogin(c)
	if err != nil {
		tools.RespUnauthorized(c)
		return
	}
	oidStr := c.PostForm("oid")
	oid, _ := strconv.Atoi(oidStr)
	err = servive.CancelOrder(id, oid)
	if err != nil {
		tools.RespFail(c, "取消订单失败")
		return
	}
	tools.RespSuccess(c, "取消订单成功")
}

func ConfirmOrder(c *gin.Context) {
	id, err := middleware.CheckLogin(c)
	if err != nil {
		tools.RespUnauthorized(c)
		return
	}
	oidStr := c.PostForm("oid")
	oid, _ := strconv.Atoi(oidStr)
	err = servive.ChangeOrder(id, oid)
	if err != nil {
		tools.RespFail(c, "确认订单状态失败")
		return
	}
	tools.RespSuccess(c, "已签收")
}
