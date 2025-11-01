package api

import "github.com/gin-gonic/gin"

func RegisterRouter(r *gin.Engine) {
	r.GET("/save", SaveUser)
	r.GET("/get", GetUserById)
	r.GET("/getall", GetAll)
	r.GET("/update", Update)
	r.GET("/delete", Delete)
}
