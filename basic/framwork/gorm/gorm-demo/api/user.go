package api

import (
	"github.com/gin-gonic/gin"
	"gorm-demo/dao"
	"time"
)

func SaveUser(c *gin.Context) {
	user := &dao.User{
		Username:   "zhangsan",
		Password:   "123456",
		CreateTime: time.Now().UnixMilli(),
	}
	dao.SaveUser(user)
	c.JSON(200, user)
}
