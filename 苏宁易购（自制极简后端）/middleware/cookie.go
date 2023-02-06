package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func CheckLogin(c *gin.Context) (uid int, err error) {
	uidstring, err := c.Cookie("user")
	uid, _ = strconv.Atoi(uidstring)
	if err != nil {
		log.Printf("the error:%#v", err)
		return
	}
	return uid, err
}
