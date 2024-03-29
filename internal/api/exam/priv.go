package exam

import (
	"context"
	"csdlpt/internal/auth"
	"csdlpt/internal/mssql"
	"csdlpt/library/ascii"
	"csdlpt/model"
	"csdlpt/pkg/token"
	"errors"
	"fmt"
	"log"
	"net/http"
)

/* */
func validateBearer(ctx context.Context, r *http.Request) (int, string, *auth.DataJWT, error) {
	var (
		excute = func(ctx context.Context, r *http.Request) (int, string, *auth.DataJWT, error) {
			var (
				// parseBearerAuth parses an HTTP Bearer Authentication string.
				// "Bearer QWxhZGRpbjpvcGVuIHNlc2FtZQ==" returns QWxhZGRpbjpvcGVuIHNlc2FtZQ.
				parseBearerAuth = func(auth string) (token string, ok bool) {
					const prefix = "Bearer "
					// Case insensitive prefix match. See Issue 22736.
					if len(auth) < len(prefix) || !ascii.EqualFold(auth[:len(prefix)], prefix) {
						return "", false
					}
					return auth[len(prefix):], true
				}
			)
			headerAuth := r.Header.Get("Authorization")
			if len(headerAuth) <= 0 {
				return http.StatusBadRequest, "", &auth.DataJWT{}, errors.New("authorization is empty")
			}
			bearer_token, ok := parseBearerAuth(headerAuth)
			if !ok {
				return http.StatusBadRequest, "", &auth.DataJWT{}, errors.New("authorization is invalid")
			}

			// get from cache DB
			// _, account_id, ok, err := fsdb.CacheLogin.Get(ctx, bearer_token)
			// if err != nil {
			// 	return http.StatusForbidden, bearer_token, &auth.DataJWT{}, err
			// }
			// if !ok {
			// 	return http.StatusForbidden, bearer_token, &auth.DataJWT{}, errors.New("token no login")
			// }
			jwt_data, status, err := auth.ValidateLoginJWT(ctx, bearer_token)
			if err != nil {
				println("ValidateLoginJWT:", err.Error())
			}
			switch status {
			case token.INPUT_EMPTY:
				return http.StatusForbidden, bearer_token, jwt_data, errors.New("token is empty")
			case token.ACCESS_TOKEN_INVALID:
				return http.StatusForbidden, bearer_token, jwt_data, errors.New("token is invalid")
			case token.ACCESS_TOKEN_EXPIRED:
				return http.StatusForbidden, bearer_token, jwt_data, errors.New("token is expired")
			case token.SUCCEED:
				// auth pass
				return http.StatusOK, bearer_token, jwt_data, nil
			default:
				return http.StatusForbidden, bearer_token, jwt_data, errors.New("validate token exception")
			}
		}
	)

	status, token, data, err := excute(ctx, r)
	if err != nil {
		println("[AUTH] ", r.RequestURI, "| Error:", err.Error())
	}
	println("[AUTH] ", r.RequestURI, "| Status:", status)
	println("[AUTH] ", r.RequestURI, "| Token:", token)
	println("[AUTH] ", r.RequestURI, "| Data:", data.UserName)
	println("[AUTH] ", r.RequestURI, "| Access Rights:", fmt.Sprintf("%+v", data))
	return status, token, data, err
}

// __getQuestionFilter
/* */
func __getQuestionFilter(ctx context.Context, request *getQuestionFilterRequest) (list *getQuestionFilterResponse, err error) {
	var (
		list_question = make([]question_data, 0)
	)
	db_questions, err := mssql.QuestionDBC.GetByStaffCode(ctx, withDBPermit(request.permit), "string")
	if err != nil {
		return &getQuestionFilterResponse{
			Code:    model.StatusServiceUnavailable,
			Message: err.Error()}, err

	}
	// if data_not_exist {
	// 	return &getQuestionFilterResponse{
	// 		Code:    model.StatusDataNotFound,
	// 		Message: "DATA_NOT_EXIST",
	// 	}, nil
	// }
	for _, question := range db_questions {
		list_question = append(list_question, withQuestionModel(&question))
	}
	return &getQuestionFilterResponse{
		Payload: filter_question_resp{
			Total:        len(list_question),
			ListQuestion: list_question,
		},
	}, nil
}

/* */
func __createQuestion(ctx context.Context, request *createQuestionRequest) (resp *createQuestionResponse, err error) {
	if request.Role != "GIANGVIEN" {
		return &createQuestionResponse{
			Code:    model.StatusForbidden,
			Message: "NOT_PERMISSION_STAFF"}, nil
	}
	id, err := ascii.GetID(request.UserName)
	if err != nil {
		return &createQuestionResponse{
			Code:    model.StatusDataNotFound,
			Message: "DATA_NOT_EXIST"}, nil
	}
	log.Println("request", request)
	var (
		question = &question_data{
			StaffCode:     id,
			Content:       request.Content,
			ChooseA:       request.ChooseA,
			ChooseB:       request.ChooseB,
			ChooseC:       request.ChooseC,
			ChooseD:       request.ChooseD,
			CorrectAnswer: request.CorrectAnswer,
			Level:         request.Level,
			CourseCode:    request.CourseCode,
		}
	)
	log.Println("question", question)
	return &createQuestionResponse{}, nil
}

/* */
func __getLastestExam(ctx context.Context, request *getLastestExamRequest) (resp *getLastestExamResponse, err error) {
	// if request.Role != "SINHVIEN" {
	// 	return &getLastestExamResponse{
	// 		Code:    model.StatusForbidden,
	// 		Message: "NOT_PERMISSION_STUDENT"}, nil
	// }

	// id, err := ascii.GetID(request.UserName)
	// if err != nil {
	// 	return &getLastestExamResponse{
	// 		Code:    model.StatusDataNotFound,
	// 		Message: "DATA_NOT_EXIST"}, nil
	// }

	var (
		list_exam = make([]exam_data, 0)
	)

	// mock_data
	for _, ele := range ExamList {
		list_exam = append(list_exam, ele)
	}

	return &getLastestExamResponse{
		Payload: filter_lastest_exam{
			Total:           len(list_exam),
			ListLastestExam: list_exam,
		},
	}, nil
}
