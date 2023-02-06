package api

import (
	"github.com/gin-gonic/gin"
	"holidy/awesomeProject/middleware"
	"holidy/awesomeProject/servive"
	"holidy/awesomeProject/tools"
	"strconv"
)

func comment(c *gin.Context) {
	id, err := middleware.CheckLogin(c)
	if err != nil {
		tools.RespUnauthorized(c)
		return
	}
	strGid := c.PostForm("gid")
	content := c.PostForm("content")
	gid, _ := strconv.Atoi(strGid)
	username := servive.GetUserById(id).Username
	err = servive.CreateComment(gid, username, content)
	if err != nil {
		tools.RespFail(c, "评论失败")
	}
	err = servive.AddComment(gid)
	if err != nil {
		tools.RespFail(c, "商品评论数增加失败")
	}
	tools.RespSuccess(c, "评论成功")
}
func viewComment(c *gin.Context) {
	_, err := middleware.CheckLogin(c)
	if err != nil {
		tools.RespUnauthorized(c)
		return
	}
	gidStr := c.PostForm("gid")
	gid, _ := strconv.Atoi(gidStr)
	comments, err := servive.GetComments(gid)
	if err != nil {
		tools.RespFail(c, "未查询到相关评论")
		return
	}
	c.JSON(200, comments)
}
