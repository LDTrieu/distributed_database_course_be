package exam

import (
	"csdlpt/internal/mssql"
	"time"
)

// type traceField struct {
// 	RequestId string `json:"reqId"`
// }

type permit struct {
	UserName   string `json:"userName"`
	FullName   string `json:"fullName"`
	PrivKey    string `json:"privKey"`
	CenterName string `json:"centerName"`
	Role       string `json:"role"`
}

func withDBPermit(p permit) mssql.DBPermitModel {
	return mssql.DBPermitModel{
		UserName:   p.UserName,
		CenterName: p.CenterName,
		PrivKey:    p.PrivKey,
	}
}

/* */
type getQuestionFilterRequest struct {
	permit
	PayLoad filter_question_req `json:"payload"`
}
type filter_question_req struct {
	CourseCode string `json:"courseCode"`
	Level      string `json:"level"`
	Size       int    `json:"size"`
}

type getQuestionFilterResponse struct {
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	Payload filter_question_resp `json:"payload"`
}

type filter_question_resp struct {
	Total        int `json:"total"`
	ListQuestion []question_data
}
type question_data struct {
	QuestionCode  string `json:"questionCode"`
	Content       string `json:"content"`
	ChooseA       string `json:"chooseA"`
	ChooseB       string `json:"chooseB"`
	ChooseC       string `json:"chooseC"`
	ChooseD       string `json:"chooseD"`
	CourseCode    string `json:"courseCode"`
	StaffCode     string `json:"staffCode"`
	Level         string `json:"level"`
	CorrectAnswer string `json:"correctAnswer"`
}

func withQuestionModel(qm *mssql.QuestionModel) question_data {
	return question_data{
		QuestionCode: qm.MaCH,
		ChooseA:      qm.ChooseA,
		ChooseB:      qm.ChooseB,
		ChooseC:      qm.ChooseC,
		ChooseD:      qm.ChooseD,
	}
}

/* */
// createQuestionRequest
type createQuestionRequest struct {
	permit
	CourseCode    string `json:"courseCode"`
	Level         string `json:"level"`
	CorrectAnswer string `json:"correctAnswer"`
	Content       string `json:"content"`
	ChooseA       string `json:"chooseA"`
	ChooseB       string `json:"chooseB"`
	ChooseC       string `json:"chooseC"`
	ChooseD       string `json:"chooseD"`
}

type createQuestionResponse struct {
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	Payload create_question_resp `json:"payload"`
}

type create_question_resp struct {
}

/* */
type getLastestExamRequest struct {
	permit
}

type getLastestExamResponse struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Payload filter_lastest_exam `json:"payload"`
}

type filter_lastest_exam struct {
	Total           int         `json:"total"`
	ListLastestExam []exam_data `json:"listLastestExam"`
}

type exam_data struct {
	ID            string    `json:"id"`
	CourseCode    string    `json:"courseCode"`
	ClassCode     string    `json:"classCode"`
	Level         string    `json:"level"`
	STExam        int       `json:"stExam"`
	TotalQuestion int       `json:"totalQuestion"`
	Duration      int       `json:"duration"`
	ExamDay       time.Time `json:"examDay"`

	Audio        string    `json:"audio"`
	ExamSeriesId int       `json:"examSeriesId"`
	Hashtag      []string  `json:"hashtag"`
	Name         string    `json:"name"`
	TotalPart    int       `json:"totalPart"`
	TotalComment int       `json:"totalComment"`
	PointReward  int       `json:"pointReward"`
	NumsJoin     int       `json:"numsJoin"`
	CreatedAt    time.Time `json:"createdAt"`
}
