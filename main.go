package main

import (
	"csdlpt/internal/api/login"
	"csdlpt/internal/api/portal"
	"csdlpt/internal/api/testpractice"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)
	login.Reg(router)
	portal.Reg(router)
	testpractice.Reg((router))

	router.Run(":8080")
}
