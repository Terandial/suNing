package api

import (
	"github.com/gin-gonic/gin"
	"holidy/awesomeProject/middleware"
	"holidy/awesomeProject/servive"
	"holidy/awesomeProject/tools"
	"strconv"
)

func charge(c *gin.Context) {
	id, err := middleware.CheckLogin(c)
	if err != nil {
		tools.RespUnauthorized(c)
		return
	}
	strMoney := c.PostForm("money")
	money, _ := strconv.ParseFloat(strMoney, 64)
	err = servive.AddMoney(id, money)
	if err != nil {
		tools.RespFail(c, "充值失败")
		return
	}
	tools.RespSuccess(c, "充值成功")
}
func view(c *gin.Context) {
	id, err := middleware.CheckLogin(c)
	if err != nil {
		tools.RespUnauthorized(c)
		return
	}
	w, err := servive.GetWallet(id)
	if err != nil {
		tools.RespFail(c, "未找到钱包")
		return
	}
	money := w.Money
	strMoney := strconv.FormatFloat(money, 'f', 2, 64)
	tools.RespSuccess(c, "钱包余额为:"+strMoney)
}
