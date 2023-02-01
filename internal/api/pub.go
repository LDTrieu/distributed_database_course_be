package api

import (
	"csdlpt/pkg/wlog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Reg(router *gin.Engine) {
	router.GET("/api/login/info", loginInfo)
	router.POST("/api/login/login", login)
	router.GET("/api/portal/ping-db", pingDB)
	router.POST("/api/portal/pong", pong)

	//Staff
	router.GET("/api/portal/list/staff", listStaff)
	router.POST("/api/portal/create/staff", createStaff)

	// Faculty
	router.GET("/api/portal/list/faculty", listFaculty)
	router.POST("/api/portal/create/faculty", createFaculty)

	// Student
	//router.GET("/api/portal/list/student", listStudent)

	// Class

	// Subject

}

/* */
func pingDB(c *gin.Context) {

	resp, err := __pingDB(c.Request.Context())
	if err != nil {
		wlog.Error(c, err)
	}
	c.JSON(http.StatusOK, resp)
}

/* */
func pong(c *gin.Context) {
	// validate token
	status, _, token_data, err := validateBearer(c.Request.Context(), c.Request)
	if err != nil {
		c.AbortWithError(status, err)
		return
	}

	var (
		request = &pongRequest{
			Permit: token_data.UserName,
		}
	)
	if err := c.BindJSON(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	resp, err := __pong(c.Request.Context(), request)
	if err != nil {
		wlog.Error(c, err)
	}

	c.JSON(http.StatusOK, resp)
}

/* */
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
func listStaff(c *gin.Context) {
	status, _, data, err := validateBearer(c.Request.Context(), c.Request)
	if err != nil {
		c.AbortWithError(status, err)
		return
	}

	var (
		request = listStaffRequest{
			permit: permit{
				UserName:   data.UserName,
				FullName:   data.FullName,
				CenterName: data.CenterName,
				Role:       data.Role,
			},
		}
	)
	resp, err := __listStaff(c.Request.Context(), &request)
	if err != nil {
		wlog.Error(c, err)
	}

	c.JSON(http.StatusOK, resp)

}

/* */

func createStaff(c *gin.Context) {
	status, _, data, err := validateBearer(c.Request.Context(), c.Request)
	if err != nil {
		c.AbortWithError(status, err)
		return
	}
	var (
		request = createStaffRequest{
			permit: permit{
				UserName:   data.UserName,
				FullName:   data.FullName,
				CenterName: data.CenterName,
				Role:       data.Role,
			},
		}
	)

	if err := c.BindJSON(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resp, err := __createStaff(c.Request.Context(), request)
	if err != nil {
		wlog.Error(c, err)
	}

	c.JSON(http.StatusOK, resp)

}

/* */

func listFaculty(c *gin.Context) {
	status, _, data, err := validateBearer(c.Request.Context(), c.Request)
	if err != nil {
		c.AbortWithError(status, err)
		return
	}

	var (
		request = listFacultyRequest{
			permit: permit{
				UserName:   data.UserName,
				FullName:   data.FullName,
				CenterName: data.CenterName,
				Role:       data.Role,
			},
		}
	)

	resp, err := __listFaculty(c.Request.Context(), &request)
	if err != nil {
		wlog.Error(c, err)
	}

	c.JSON(http.StatusOK, resp)

}

/* */
func createFaculty(c *gin.Context) {
	status, _, data, err := validateBearer(c.Request.Context(), c.Request)
	if err != nil {
		c.AbortWithError(status, err)
		return
	}
	var (
		request = createStaffRequest{
			permit: permit{
				UserName:   data.UserName,
				FullName:   data.FullName,
				CenterName: data.CenterName,
				Role:       data.Role,
			},
		}
	)

	if err := c.BindJSON(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resp, err := __createStaff(c.Request.Context(), request)
	if err != nil {
		wlog.Error(c, err)
	}

	c.JSON(http.StatusOK, resp)

}
