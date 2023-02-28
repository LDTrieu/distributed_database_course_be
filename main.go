package main

import (
	"csdlpt/internal/api/exam"
	"csdlpt/internal/api/login"
	"csdlpt/internal/api/portal"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.Use(cors.Default())

	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://localhost:8080"}                   // Các origin được phép truy cập API
	// config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"} // Các method được phép truy cập
	// config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"} // Các header được phép truy cập
	// config.ExposeHeaders = []string{"Content-Length"}                         // Các header sẽ được expose

	login.Reg(router)
	portal.Reg(router)
	exam.Reg((router))

	//router.Use(cors.New(config))
	router.Run(":8080")
}
