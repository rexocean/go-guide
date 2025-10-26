package main

import (
	"github.com/gin-gonic/gin"
	"gorm-demo/router"
	"log"
)

func main() {
	r := gin.Default()
	err := r.Run(":8080")
	router.InitRouter(r)
	if err != nil {
		log.Fatal(err)
	}

}
