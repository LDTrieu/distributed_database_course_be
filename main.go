package main

import (
	"csdlpt/internal/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)
	api.Reg(router)

	router.Run(":8080")
}
