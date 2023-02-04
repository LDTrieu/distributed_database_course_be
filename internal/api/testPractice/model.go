package testpractice

import "csdlpt/internal/mssql"

type traceField struct {
	RequestId string `json:"reqId"`
}

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
	QuestionCode string `json:"questionCode"`
	Content      string `json:"content"`
	ChooseA      string `json:"chooseA"`
	ChooseB      string `json:"chooseB"`
	ChooseC      string `json:"chooseC"`
	ChooseD      string `json:"chooseD"`
}
type filter_question_req struct {
	CourseCode string `json:"courseCode"`
	Level      string `json:"level"`
	Size       int    `json:"size"`
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
