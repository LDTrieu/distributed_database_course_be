package login

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

	"github.com/google/uuid"
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

/* */
// Get DS Phan Manh
func __loginInfo(ctx context.Context) (*loginInfoResponse, error) {
	// get list publisher_name from DB
	db_pubs, err := mssql.DBServerDBC.GetListCenter(ctx)
	if err != nil {
		log.Println("ERR: ", err)
		return &loginInfoResponse{
			Code:    model.StatusServiceUnavailable,
			Message: err.Error()}, err

	}
	//log.Println("db_pubs: ", db_pubs)
	// DB layer to Handle layer
	var (
		list_center = make([]string, 0)
	)

	for _, center := range db_pubs {
		list_center = append(list_center, string(center))
	}

	return &loginInfoResponse{
		Payload: login_info_resp{
			ListCenter:  list_center,
			TotalCenter: len(list_center),
		},
	}, nil
}

/* */

func __login(ctx context.Context, request *loginRequest) (*loginResponse, error) {
	// Validate request
	if err := request.validate(); err != nil {
		return &loginResponse{
			Code:    model.StatusBadRequest,
			Message: "DATA_INVALID",
		}, nil
	}
	log.Println(request)

	_, ho_ten, first_name, last_name, staff_role, data_exist, err := mssql.DBServerDBC.Login(ctx, withRequestPermission(request), request.UserName)
	if err != nil {
		return &loginResponse{
			Code:    model.StatusServiceUnavailable,
			Message: err.Error(),
		}, nil
	}
	if !data_exist {
		return &loginResponse{
			Code:    model.StatusDataNotFound,
			Message: "DATA_NOT_EXIST",
		}, nil
	}
	//validateBearer(ctx, ctx.Req)
	// Gen auth token
	_, jwt_login, err := auth.GenerateJWTLoginSession(ctx, request.UserName, ho_ten, staff_role, request.CenterName, uuid.New().String())
	if err != nil {
		return &loginResponse{
			Code:    model.StatusServiceUnavailable,
			Message: err.Error(),
		}, nil
	}

	return &loginResponse{
		Payload: login_resp{
			UserName:  request.UserName,
			FirstName: first_name,
			LastName:  last_name,
			Token:     jwt_login.Token,
		},
	}, nil
}

/* */
func __getMe(ctx context.Context, request *getUserMeRequest) (*getUserMeResponse, error) {

	return &getUserMeResponse{
		Payload: info_user_me{
			UserName:    request.UserName,
			RoleName:    request.Role,
			FirstName:   "FirstName",
			LastName:    "LastName",
			PhoneNumber: "0948518286",
			Avt:         "https://media.istockphoto.com/id/1223671392/vector/default-profile-picture-avatar-photo-placeholder-vector-illustration.jpg?s=170667a&w=0&k=20&c=m-F9Doa2ecNYEEjeplkFCmZBlc5tm1pl1F7cBCh9ZzM=",
			Banner:      "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAX8AAACECAMAAABPuNs7AAAACVBMVEWAgICLi4uUlJSuV9pqAAABI0lEQVR4nO3QMQEAAAjAILV/aGPwjAjMbZybnTjbP9W/1b/Vv9W/1b/Vv9W/1b/Vv9W/1b/Vv9W/1b/Vv9W/1b/Vv9W/1b/Vv9W/1b/Vv9W/1b/Vv9W/1b/Vv9W/1b/Vv9W/1b/Vv9W/1b/Vv9W/1b/Vv9W/1b/Vv9W/1b/Vv9W/1b/Vv9W/1b/Vv9W/1b/Vv9W/1b/Vv9W/1b/Vv9W/1b/Vv9W/1b/Vv9W/1b/Vv9W/1b/Vv9W/1b/Vv9W/1b/Vv9W/1b/Vv9W/1b/Vv9W/1b/Vv9W/1b/Vv9W/1b/Vv9W/1b/Vv9W/1b/Vv9W/1b/Vv9W/1b/Vv9W/1b/Vv9W/1b/Vv9W/1b/Vv9W/1b/Vv9W/1b/Vv9W/1b/Vv9W/1b+1cxvnHi9hBAfkOyqGAAAAAElFTkSuQmCC",
			UserId:      "123",
		},
	}, nil
}
