package main

import (
	"csdlpt/internal/api/login"
	"csdlpt/internal/api/portal"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)
	login.Reg(router)
	portal.Reg(router)

	router.Run(":8080")
}
