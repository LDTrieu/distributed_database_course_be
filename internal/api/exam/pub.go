package exam

import (
	"csdlpt/pkg/wlog"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Reg(router *gin.Engine) {

	//
	router.POST("/api/test/get-filter", getQuestionFilter)

	// Giang Vien Dang Ky - Tạo kỳ thi
	// router.GET("/api/portal/list/exam", listExam)

	// Bo De - Câu hỏi
	router.POST("/api/portal/create/question", createQuestion)

	// Thi
	router.GET("/api/exam/list/latest-exam", getLastestExam)     // Latest exam
	router.GET("/api/exam/list/exam-takingg/:id", getTakingExam) // Full Exam

}

// getQuestionFilter
func getQuestionFilter(c *gin.Context) {
	status, _, data, err := validateBearer(c.Request.Context(), c.Request)
	if err != nil {
		c.AbortWithError(status, err)
		return
	}
	var (
		request = getQuestionFilterRequest{
			permit: permit{
				UserName:   data.UserName,
				FullName:   data.FullName,
				CenterName: data.CenterName,
				Role:       data.Role,
			},
		}
	)
	resp, err := __getQuestionFilter(c.Request.Context(), &request)
	if err != nil {
		wlog.Error(c, err)
	}
	c.JSON(http.StatusOK, resp)
}

// createQuestion
func createQuestion(c *gin.Context) {
	status, _, data, err := validateBearer(c.Request.Context(), c.Request)
	if err != nil {
		c.AbortWithError(status, err)
		return
	}

	var (
		request = createQuestionRequest{
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

	resp, err := __createQuestion(c.Request.Context(), &request)
	if err != nil {
		wlog.Error(c, err)
	}
	c.JSON(http.StatusOK, resp)
}

// getLastestExam
func getLastestExam(c *gin.Context) {
	status, _, data, err := validateBearer(c.Request.Context(), c.Request)
	if err != nil {
		c.AbortWithError(status, err)
		return
	}

	var (
		request = getLastestExamRequest{
			permit: permit{
				UserName:   data.UserName,
				FullName:   data.FullName,
				CenterName: data.CenterName,
				Role:       data.Role,
			},
		}
	)

	resp, err := __getLastestExam(c.Request.Context(), &request)
	if err != nil {
		wlog.Error(c, err)
	}
	c.JSON(http.StatusOK, resp)
}

/* */
func getTakingExam(c *gin.Context) {
	status, _, data, err := validateBearer(c.Request.Context(), c.Request)
	if err != nil {
		c.AbortWithError(status, err)
		return
	}

	var (
		request = getTakingExamRequest{
			permit: permit{
				UserName:   data.UserName,
				FullName:   data.FullName,
				CenterName: data.CenterName,
				Role:       data.Role,
			},
			ExamId: c.Param("id"),
		}
	)
	log.Println("ExamId: ", request.ExamId)

	resp, err := __getTakingExam(c.Request.Context(), &request)
	if err != nil {
		wlog.Error(c, err)
	}
	c.JSON(http.StatusOK, resp)
}
