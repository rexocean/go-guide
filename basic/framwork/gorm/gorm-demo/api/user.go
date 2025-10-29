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

func GetUserById(c *gin.Context) {
	user := dao.GetById(1)
	c.JSON(200, user)
}

func GetAll(c *gin.Context) {
	users := dao.GetAll()
	c.JSON(200, users)
}

func Update(c *gin.Context) {
	dao.Update(1)
	user := dao.GetById(1)
	c.JSON(200, user)
}

func Delete(c *gin.Context) {
	dao.Delete(1)
	user := dao.GetById(1)
	c.JSON(200, user)
}
