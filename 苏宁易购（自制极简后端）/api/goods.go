package api

import (
	"github.com/gin-gonic/gin"
	"holidy/awesomeProject/middleware"
	"holidy/awesomeProject/servive"
	"holidy/awesomeProject/tools"
	"strconv"
)

func allProduct(c *gin.Context) {
	_, err := middleware.CheckLogin(c)
	if err != nil {
		tools.RespUnauthorized(c)
		return
	}
	g, err := servive.GetAllGoods()
	if err != nil {
		tools.RespFail(c, "未能获取到商品数据")
		return
	}
	c.JSON(200, g)
}

func addInCart(c *gin.Context) {
	id, err := middleware.CheckLogin(c)
	if err != nil {
		tools.RespUnauthorized(c)
		return
	}
	gidStr := c.PostForm("gid")
	gid, _ := strconv.Atoi(gidStr)
	err = servive.AddProductInCart(id, gid)
	if err != nil {
		tools.RespFail(c, "加入购物车失败")
		return
	}
	tools.RespSuccess(c, "加入购物车成功")
}

func delInCart(c *gin.Context) {
	id, err := middleware.CheckLogin(c)
	if err != nil {
		tools.RespUnauthorized(c)
		return
	}
	gidStr := c.PostForm("gid")
	gid, _ := strconv.Atoi(gidStr)
	err = servive.DelProductInCart(id, gid)
	if err != nil {
		tools.RespFail(c, "删除该商品失败")
		return
	}
	tools.RespSuccess(c, "删除该商品成功")
}
func WatchCart(c *gin.Context) {
	id, err := middleware.CheckLogin(c)
	if err != nil {
		tools.RespUnauthorized(c)
		return
	}
	cart, err := servive.GetAllProductInCart(id)
	if err != nil {
		tools.RespFail(c, "未获取到购物车内商品")
		return
	}
	c.JSON(200, cart)
}
