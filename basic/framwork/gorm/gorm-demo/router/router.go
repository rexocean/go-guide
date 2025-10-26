package router

import (
	"github.com/gin-gonic/gin"
	"gorm-demo/api"
)

func InitRouter(engine *gin.Engine) {
	api.RegisterRouter(engine)
}
