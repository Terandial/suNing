package api

import (
	"github.com/gin-gonic/gin"
	"holidy/awesomeProject/middleware"
	"holidy/awesomeProject/servive"
	"holidy/awesomeProject/tools"
	"strconv"
	"strings"
)

func buyFromCart(c *gin.Context) {
	id, err := middleware.CheckLogin(c)
	if err != nil {
		tools.RespUnauthorized(c)
		return
	}

	//前端输入商品gid时需要用逗号间隔
	gidStr := c.PostForm("gid")
	aidStr := c.PostForm("aid")
	aid, _ := strconv.Atoi(aidStr)
	gidSlice := strings.Split(gidStr, ",")
	count := len(gidSlice)
	gid := make([]int, count)
	for i := 0; i < count; i++ {
		a, _ := strconv.Atoi(gidSlice[i])
		gid = append(gid, a)
	}
	need, err := servive.ComputePrice(id, gid)
	if err != nil {
		tools.RespFail(c, "计算总金额失败")
		return
	}
	flag, err := servive.MinusMoney(id, need)
	if err != nil {
		tools.RespFail(c, "扣除余额失败")
		return
	}
	if flag {
		tools.RespFail(c, "您的余额不足")
		return
	}
	address, err := servive.SearchAddressByAid(aid)
	receipt := address.Place
	for i := 0; i < count; i++ {
		good, err := servive.SearchGoodByGid(gid[i])
		if err != nil {
			tools.RespFail(c, "获取商品失败")
			return
		}
		gname := good.GName
		sid := good.Sid
		store, err := servive.SearchStoreBySid(sid)
		if err != nil {
			tools.RespFail(c, "获取商店信息失败")
			return
		}
		shipment := store.Address
		money, err := servive.ComputeOrderPrice(id, gid[i])
		if err != nil {
			tools.RespFail(c, "计算订单金额失败")
			return
		}
		err = servive.CreateOrder(id, gid[i], gname, shipment, receipt, money)
		if err != nil {
			tools.RespFail(c, "生成订单失败")
			return
		}
	}
	tools.RespSuccess(c, "购买成功，订单已生成")
}

func buyInShop(c *gin.Context) {
	id, err := middleware.CheckLogin(c)
	if err != nil {
		tools.RespUnauthorized(c)
		return
	}
	gidStr := c.PostForm("gid")
	number := c.PostForm("number")
	aidStr := c.PostForm("aid")
	aid, _ := strconv.Atoi(aidStr)
	gid, _ := strconv.Atoi(gidStr)
	g, err := servive.SearchGoodByGid(gid)
	if err != nil {
		tools.RespFail(c, "未找到该商品")
		return
	}
	need, err := servive.ComputeShopPrice(gid, number)
	if err != nil {
		tools.RespFail(c, "计算金额失败")
		return
	}
	flag, err := servive.MinusMoney(id, need)
	if err != nil {
		tools.RespFail(c, "扣除余额失败")
		return
	}
	if flag {
		tools.RespFail(c, "您的余额不足")
		return
	}
	address, err := servive.SearchAddressByAid(aid)
	receipt := address.Place
	sid := g.Sid
	gname := g.GName
	s, err := servive.SearchStoreBySid(sid)
	if err != nil {
		tools.RespFail(c, "获取商品所属商店信息失败")
	}
	shipment := s.Address
	err = servive.CreateOrder(id, gid, gname, shipment, receipt, need)
	if err != nil {
		tools.RespFail(c, "生成订单失败")
		return
	}
	tools.RespSuccess(c, "生成订单成功")
}
