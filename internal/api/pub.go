package api

import (
	"csdlpt/pkg/wlog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Reg(router *gin.Engine) {
	router.GET("/api/portal/ping-db", pingDB)
	router.GET("/api/login/info", loginInfo)
	router.POST("/api/login/login", login)
}

func pingDB(c *gin.Context) {

	resp, err := __pingDB(c.Request.Context())
	if err != nil {
		// log it
		// c.JSON(http.StatusOK, resp)
		// return
	}
	c.JSON(http.StatusOK, resp)
}

/* */
func loginInfo(c *gin.Context) {

	resp, err := __loginInfo(c.Request.Context())
	if err != nil {
		//wlog.Error(c, err)
	}

	c.JSON(http.StatusOK, resp)

}

/* */
func login(c *gin.Context) {
	var (
		request = loginRequest{}
	)
	if err := c.BindJSON(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	resp, err := __login(c.Request.Context(), &request)
	if err != nil {
		wlog.Error(c, err)
	}

	c.JSON(http.StatusOK, resp)
}
