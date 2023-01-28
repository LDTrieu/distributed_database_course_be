package api

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
func __pingDB(ctx context.Context) (*pingDBResponse, error) {
	// get count(*) table
	a, err := mssql.DBServerDBC.Ping(ctx)
	if err != nil {
		log.Println("ERR: ", err)
	}
	log.Println(a)

	// get list table
	return nil, nil

}

/* */
func __pong(ctx context.Context, request *pongRequest) (*pongResponse, error) {
	// handle
	return &pongResponse{
		Payload: pong_resp{
			UserName: request.Permit,
		},
	}, nil
}

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
			log.Println("jwt_data.UserName", jwt_data.UserName)
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
	_, ho_ten, _, data_exist, err := mssql.DBServerDBC.Login(ctx, request.UserName)
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
	_, jwt_login, err := auth.GenerateJWTLoginSession(ctx, request.UserName, ho_ten, request.Role, request.CenterName, uuid.New().String())
	if err != nil {
		return &loginResponse{
			Code:    model.StatusServiceUnavailable,
			Message: err.Error(),
		}, nil
	}
	// // exist user_name , password != nil => pass
	// switch request.Role {
	// case "giang_vien": // role is giang_vien
	// 	log.Println("giang_vien")

	// 	user_name, err := mssql.DBServerDBC.GetListCenter(ctx)
	// 	if err != nil {
	// 		log.Println("ERR: ", err)
	// 		return &loginResponse{
	// 			Code:    model.StatusServiceUnavailable,
	// 			Message: err.Error()}, err

	// 	}
	// case "sinh_vien": // role is sinh_vien
	// 	log.Println("sinh_vien")
	// case "co_so": // role is co_so
	// 	log.Println("co_so")
	// case "truong": // role is truong
	// 	log.Println("truong")
	// default:
	// 	return &loginResponse{
	// 		Code:    model.StatusDataNotFound,
	// 		Message: "DATA_NOT_FOUND",
	// 	}, nil
	// }

	//log.Println("db_pubs: ", db_pubs)
	// DB layer to Handle layer

	// call SP_DANG_NHAP

	return &loginResponse{
		Payload: login_resp{
			UserName: request.UserName,
			Token:    jwt_login.Token,
		},
	}, nil
}

/* */
func __listStaff(ctx context.Context, request *listStaffRequest) (list *listStaffResponse, err error) {
	var (
		list_staff = make([]staff_data, 0)
		db_staffs  = make([]mssql.StaffModel, 0)
	)
	// Check center

	db_staffs, err = mssql.StaffDBC.GetAll(ctx, withDBPermit(request.permit))
	if err != nil {
		return &listStaffResponse{
			Code:    model.StatusServiceUnavailable,
			Message: err.Error()}, err

	}
	for _, staff := range db_staffs {
		list_staff = append(list_staff, withStaffModel(&staff))
	}
	return &listStaffResponse{
		Payload: list_staff_resp{
			TotalStaff: len(list_staff),
			ListStaff:  list_staff,
		},
	}, nil
}

func __createStaff(ctx context.Context, request createStaffRequest) (*createStaffResponse, error) {

	// Check Permission and StaffRole
	switch request.Role {
	case "TRUONG":
		log.Println("TRUONG")
		// StaffRole != TRUONG -> return
		if request.StaffRole != "TRUONG" {
			return &createStaffResponse{
				Code:    model.StatusForbidden,
				Message: "NOT_PERMISSION_UNI"}, nil
		}
		// RUN
		// check data_exist (Check Mã Giảng Viên)

		// run SP tạo đăng nhập

	case "COSO":
		log.Println("COSO")
		switch request.StaffRole {
		case "GIANGVIEN":
			log.Println("GIANGVIEN")

		case "COSO":
			log.Println("COSO")

		default:
			return &createStaffResponse{
				Code:    model.StatusForbidden,
				Message: "NOT_PERMISSION_CENTER"}, nil
		}
	case "GIANGVIEN":
		log.Println("GIANGVIEN")
		return &createStaffResponse{
			Code:    model.StatusForbidden,
			Message: "NOT_PERMISSION_CENTER"}, nil
	default:
		log.Println("ERROR")
		return &createStaffResponse{
			Code:    model.StatusDataNotFound,
			Message: "DATA_NOT_FOUND"}, nil
	}

	// Check data_exist in DB
	data_staff, data_exist, err := mssql.StaffDBC.Get(ctx, withDBPermit(request.permit), request.UserName)
	if err != nil {
		return &createStaffResponse{
			Code:    model.StatusServiceUnavailable,
			Message: err.Error()}, err

	}
	if data_exist {
		return &createStaffResponse{
			Code:    model.StatusDataDuplicated,
			Message: err.Error()}, errors.New("resource already exists")
	}
	log.Println(data_staff.TenNhom)

	// add item to DB
	// create account with Staff Permision

	return &createStaffResponse{}, nil
}

/* */
func __listFaculty(ctx context.Context, request *listFacultyRequest) (list *listFacultyResponse, err error) {
	var (
		list_faculty = make([]faculty_data, 0)
		db_facultys  = make([]mssql.FacultyModel, 0)
	)
	// Check center

	db_facultys, err = mssql.FacultyDBC.GetAll(ctx, withDBPermit(request.permit))
	if err != nil {
		return &listFacultyResponse{
			Code:    model.StatusServiceUnavailable,
			Message: err.Error()}, err

	}
	for _, faculty := range db_facultys {
		list_faculty = append(list_faculty, withFacultyModel(&faculty))
	}
	return &listFacultyResponse{
		Payload: list_faculty_resp{
			TotalFaculty: len(list_faculty),
			ListFaculty:  list_faculty,
		},
	}, nil
}

// /* */
// func __createCenterStaff(ctx context.Context, request createCenterStaffRequest) (*createCenterStaffResponse, error) {

// 	// Check permition request - Role Center
// 	if request.permit.Role != "CENTER" {
// 		return &createCenterStaffResponse{
// 				Code:    model.StatusForbidden,
// 				Message: "ACCESS_DENIED"},
// 			errors.New("user access denied")
// 	}
// 	// Check data_exist in DB
// 	_, data_exist, err := mssql.CenterStaffDBC.Get(ctx, withDBPermit(request.permit), request.UserName)
// 	if err != nil {
// 		return &createCenterStaffResponse{
// 			Code:    model.StatusServiceUnavailable,
// 			Message: err.Error()}, err

// 	}
// 	if data_exist {
// 		return &createCenterStaffResponse{
// 			Code:    model.StatusDataDuplicated,
// 			Message: err.Error()}, errors.New("resource already exists")
// 	}

// 	// add item to DB
// 	// create account with Staff Permision

// 	return &createCenterStaffResponse{}, nil
// }
