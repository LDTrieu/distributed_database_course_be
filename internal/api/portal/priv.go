package portal

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
	if data_not_exist {
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

/* */
func __createClass(ctx context.Context, request createClassRequest) (*createClassResponse, error) {
	var (
		class = &class_data{
			ClassCode:   request.ClassCode,
			ClassName:   request.ClassName,
			FacultyCode: request.FacultyCode,
		}
	)

	// Check Permission and StaffRole
	switch request.Role {
	case "COSO":
		if err := mssql.ClassDBC.Create(ctx, withDBPermit(request.permit), withClassData(class)); err != nil {
			return &createClassResponse{
				Code:    model.StatusServiceUnavailable,
				Message: err.Error()}, err
		}
	default:
		return &createClassResponse{
			Code:    model.StatusForbidden,
			Message: "NOT_PERMISSION"}, nil
	}

	return &createClassResponse{}, nil
}

/* */
func __listCourse(ctx context.Context, request *listCourseRequest) (list *listCourseResponse, err error) {
	var (
		list_course = make([]course_data, 0)
		//db_classes = make([]mssql.ClassModel, 0)
	)
	db_courses, data_not_exist, err := mssql.CourseDBC.GetAll(ctx, withDBPermit(request.permit))
	if err != nil {
		return &listCourseResponse{
			Code:    model.StatusServiceUnavailable,
			Message: err.Error()}, err

	}
	if data_not_exist {
		return &listCourseResponse{
			Code:    model.StatusDataNotFound,
			Message: "DATA_NOT_EXIST",
		}, nil
	}
	for _, course := range db_courses {
		list_course = append(list_course, withCourseModel(&course))
	}
	return &listCourseResponse{
		Payload: list_course_resp{
			TotalCourse: len(list_course),
			ListCourse:  list_course,
		},
	}, nil
}

/* */
func __listMockCourse(ctx context.Context, request *listMockCourseRequest) (list *listMockCourseResponse, err error) {
	var (
		list_mock_course = make([]mock_course_data, 0)
		//db_classes = make([]mssql.ClassModel, 0)
	)
	// db_courses, data_not_exist, err := mssql.CourseDBC.GetAll(ctx, withDBPermit(request.permit))
	// if err != nil {
	// 	return &listCourseResponse{
	// 		Code:    model.StatusServiceUnavailable,
	// 		Message: err.Error()}, err

	// }
	// if data_not_exist {
	// 	return &listCourseResponse{
	// 		Code:    model.StatusDataNotFound,
	// 		Message: "DATA_NOT_EXIST",
	// 	}, nil
	// }
	// for _, course := range db_courses {
	// 	list_course = append(list_mock_course, withCourseModel(&course))
	// }
	ele := mock_course_data{
		CourseCode:     "1",
		CourseName:     "Khóa nhập môn TOEIC - Các kiến thức nền tảng",
		Image:          "http://res.cloudinary.com/doxsstgkc/image/upload/v1675147193/examify/image2_efnpkj.png",
		Level:          "general",
		Charges:        true,
		PointToUnlock:  0,
		PointReward:    2000,
		QuantityRating: 0,
		AvgRating:      "0.00",
		Participants:   3,
		Price:          1000,
		Discount:       10,
		TotalChapter:   5,
		TotalLesson:    24,
		TotalVideoTime: 1,
		Achieves:       "<p>Xây dựng được một nền tảng vững chắc để đi bước đầu trên con đường học TOEIC</p><p>Thích nghi với các dạng khó hơn trong đề thi TOEIC về sau</p>",
		Description:    "<ol><li>Trợ động từ, động từ đặc biệt</li><li>Rút gọn mệnh đề quan hệ</li><li>Động từ V1 - To V1 - V-ing</li><li>Phân từ (- ed, - ing) và mệnh đề phân từ</li><li>Sự phủ định và công thức song song</li><li>Các dạng so sánh (hơn - bằng - nhất ...)</li></ol>",
		CreateBy:       1,
		CreatedAt:      "2023-01-31T06:39:53.875Z",
		UpdatedAt:      "2023-02-22T17:46:15.100Z",
		IsJoin:         true,
	}
	list_mock_course = append(list_mock_course, ele)
	return &listMockCourseResponse{
		Payload: list_mock_course_resp{
			TotalCourse: 1,
			ListCourse:  list_mock_course,
		},
	}, nil
}

/* */
func __createCourse(ctx context.Context, request createCourseRequest) (*createCourseResponse, error) {
	var (
		course = &course_data{
			CourseCode: request.CourseCode,
			CourseName: request.CourseName,
		}
	)

	// Check Permission and StaffRole
	switch request.Role {
	case "COSO":
		if err := mssql.CourseDBC.Create(ctx, withDBPermit(request.permit), withCourseData(course)); err != nil {
			return &createCourseResponse{Code: model.StatusServiceUnavailable, Message: err.Error()}, err
		}
	default:
		return &createCourseResponse{Code: model.StatusForbidden, Message: "NOT_PERMISSION"}, nil
	}

	return &createCourseResponse{}, nil
}

/* */
func __listStudent(ctx context.Context, request *listStudentRequest) (list *listStudentResponse, err error) {
	var (
		list_student = make([]student_data, 0)
		//db_classes = make([]mssql.ClassModel, 0)
	)
	db_students, data_not_exist, err := mssql.StudentDBC.GetAll(ctx, withDBPermit(request.permit))
	if err != nil {
		return &listStudentResponse{
			Code:    model.StatusServiceUnavailable,
			Message: err.Error()}, err

	}
	if data_not_exist {
		return &listStudentResponse{
			Code:    model.StatusDataNotFound,
			Message: "DATA_NOT_EXIST",
		}, nil
	}
	for _, student := range db_students {
		list_student = append(list_student, withStudentModel(&student))
	}
	return &listStudentResponse{
		Payload: list_student_resp{
			TotalStudent: len(list_student),
			ListStudent:  list_student,
		},
	}, nil
}
