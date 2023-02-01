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

	_, ho_ten, staff_role, data_exist, err := mssql.DBServerDBC.Login(ctx, withRequestPermission(request), request.UserName)
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
	var (
		// Nếu chưa, tạo tài khoản mới
		staff = &staff_data{
			UserName:    request.LoginName,
			FirstName:   request.FirstName,
			LastName:    request.LastName,
			FullName:    request.FullName,
			Address:     request.Address,
			FacultyCode: request.FacultyCode,
			StaffRole:   request.StaffRole,
		}

		create_login = func() (_ *createStaffResponse, err error) {
			// Check user_name exist
			_, data_not_exist, err := mssql.StaffDBC.GetStaffWithoutUserName(ctx, withDBPermit(request.permit), request.LoginName)
			if err != nil {
				return &createStaffResponse{
					Code:    model.StatusServiceUnavailable,
					Message: err.Error()}, err

			}
			if !data_not_exist { // Neu exist -> return err
				return &createStaffResponse{
					Code:    model.StatusDataDuplicated,
					Message: "DATA_ALREADY_EXIST"}, nil
			}
			// SP tạo GiangVien
			if err := mssql.StaffDBC.Create(ctx, withDBPermit(request.permit), withStaffData(staff)); err != nil {
				return &createStaffResponse{
					Code:    model.StatusServiceUnavailable,
					Message: err.Error()}, err
			}

			// run SP tạo đăng nhập
			if err := mssql.StaffDBC.CreateLogin(ctx, withDBPermit(request.permit), withStaffData(staff)); err != nil {
				return &createStaffResponse{
					Code:    model.StatusServiceUnavailable,
					Message: err.Error()}, err
			}
			return &createStaffResponse{}, nil
		}
	)

	// Check Permission and StaffRole
	switch request.Role {
	case "TRUONG":
		// TH20x
		if request.StaffRole != "TRUONG" {
			return &createStaffResponse{
				Code:    model.StatusForbidden,
				Message: "NOT_PERMISSION_UNI"}, nil
		}
		create_login_resp, err := create_login()
		if err != nil {
			return create_login_resp, err
		}

	case "COSO":
		switch request.StaffRole {
		case "COSO":
			// TH30x
			create_login_resp, err := create_login()
			if err != nil {
				return create_login_resp, err
			}
		case "GIANGVIEN":
			// TH50x
			create_login_resp, err := create_login()
			if err != nil {
				return create_login_resp, err
			}
		default:
			return &createStaffResponse{
				Code:    model.StatusForbidden,
				Message: "NOT_PERMISSION_CENTER"}, nil
		}
	case "GIANGVIEN":
		return &createStaffResponse{
			Code:    model.StatusForbidden,
			Message: "NOT_PERMISSION_GIANGVIEN"}, nil
	default:
		return &createStaffResponse{
			Code:    model.StatusDataNotFound,
			Message: "DATA_NOT_FOUND"}, nil
	}
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

/* */
func __createFaculty(ctx context.Context, request createFacultyRequest) (*createFacultyResponse, error) {
	var (
		faculty = &faculty_data{
			FacultyCode: request.FacultyCode,
			FacultyName: request.FacultyName,
			CenterCode:  request.CenterCode,
		}
	)

	// Check Permission and StaffRole
	switch request.Role {
	case "COSO":
		if err := mssql.FacultyDBC.Create(ctx, withDBPermit(request.permit), withFacultyData(faculty)); err != nil {
			return &createFacultyResponse{
				Code:    model.StatusServiceUnavailable,
				Message: err.Error()}, err
		}
	default:
		return &createFacultyResponse{
			Code:    model.StatusForbidden,
			Message: "NOT_PERMISSION"}, nil
	}

	return &createFacultyResponse{}, nil
}

/* */
func __listClass(ctx context.Context, request *listClassRequest) (list *listClassResponse, err error) {
	var (
		list_class = make([]class_data, 0)
		//db_classes = make([]mssql.ClassModel, 0)
	)
	db_classes, data_not_exist, err := mssql.ClassDBC.GetAll(ctx, withDBPermit(request.permit))
	if err != nil {
		return &listClassResponse{
			Code:    model.StatusServiceUnavailable,
			Message: err.Error()}, err

	}
	if !data_not_exist {
		return &listClassResponse{
			Code:    model.StatusDataNotFound,
			Message: "DATA_NOT_EXIST",
		}, nil
	}
	for _, class := range db_classes {
		list_class = append(list_class, withClassModel(&class))
	}
	return &listClassResponse{
		Payload: list_class_resp{
			TotalClass: len(list_class),
			ListClass:  list_class,
		},
	}, nil
}
