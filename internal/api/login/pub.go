package login

import (
	"csdlpt/pkg/wlog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Reg(router *gin.Engine) {
	router.GET("/api/login/info", loginInfo)
	router.POST("/api/login/login", login)
	router.GET("/api/login/me", getUserMe)

}

/* API Get CoSo*/
func loginInfo(c *gin.Context) {

	resp, err := __loginInfo(c.Request.Context())
	if err != nil {
		wlog.Error(c, err)
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

/* */
func getUserMe(c *gin.Context) {
	status, _, data, err := validateBearer(c.Request.Context(), c.Request)
	if err != nil {
		c.AbortWithError(status, err)
		return
	}
	var (
		request = getUserMeRequest{
			permit: permit{
				UserName:   data.UserName,
				FullName:   data.FullName,
				CenterName: data.CenterName,
				Role:       data.Role,
			},
		}
	)
	resp, err := __getMe(c.Request.Context(), &request)
	if err != nil {
		wlog.Error(c, err)
	}

	c.JSON(http.StatusOK, resp)
}
