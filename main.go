package main

import (
	"dydemo/model"

	"github.com/gin-gonic/gin"
)

func main() {
	model.InitDB()
	r := gin.Default()
	initRouter(r)
	r.Run(":8080")
}
