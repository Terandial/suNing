package api

import (
	"github.com/gin-gonic/gin"
	"holidy/awesomeProject/middleware"
	"holidy/awesomeProject/servive"
	"holidy/awesomeProject/tools"
	"strconv"
)

func like(c *gin.Context) {
	id, err := middleware.CheckLogin(c)
	if err != nil {
		tools.RespUnauthorized(c)
		return
	}
	gidStr := c.PostForm("gid")
	gid, err := strconv.Atoi(gidStr)
	err = servive.AddLike(id, gid)
	if err != nil {
		tools.RespFail(c, "收藏失败")
		return
	}
	tools.RespSuccess(c, "收藏成功")
}

func unlike(c *gin.Context) {
	id, err := middleware.CheckLogin(c)
	if err != nil {
		tools.RespUnauthorized(c)
		return
	}
	gidStr := c.PostForm("gid")
	gid, err := strconv.Atoi(gidStr)
	err = servive.MinusLike(id, gid)
	if err != nil {
		tools.RespFail(c, "取消收藏失败")
		return
	}
	tools.RespSuccess(c, "已取消")
}
