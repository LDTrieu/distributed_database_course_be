package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Reg(router *gin.Engine) {
	router.GET("/api/portal/ping-db", pingDB)
	router.GET("/api/portal/", pingDB)
}

func pingDB(c *gin.Context) {

	resp, err := __pingDB(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusOK, resp)
		return
	}

	// Trace client and result
	//resp.traceField = request.traceField
	c.JSON(http.StatusOK, resp)
}
