package api

import (
	"csdlpt/internal/mssql"
	"errors"
	"log"
	"time"
)

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

type pingDBResponse struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Payload ping_db_resp `json:"payload"`
}

type ping_db_resp struct {
	// DB name
	DBName string `json:"dbName"`
	// list table (string)
	ListTable string `json:"listTable"`
}

/* */

type loginInfoResponse struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Payload login_info_resp `json:"payload"`
}
type login_info_resp struct {
	TotalCenter int      `json:"totalCenter"`
	ListCenter  []string `json:"listCenter"`
}

/* */
type loginRequest struct {
	UserName   string `json:"userName"`
	Password   string `json:"password"`
	Role       string `json:"role"` // giang_vien || sinh_vien || truong || co_so
	CenterName string `json:"centerName"`
}

type loginResponse struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Payload login_resp `json:"payload"`
}

type login_resp struct {
	UserName string `json:"userName"`
	Token    string `json:"token"`
}

func (ins *loginRequest) validate() error {
	if len(ins.UserName) < 1 {
		log.Println("ins.UserName", ins.UserName)
		return errors.New("field UserName is invalid")
	}
	if len(ins.Password) < 1 {
		log.Println("ins.Password", ins.UserName)
		return errors.New("field Password is invalid")
	}
	if len(ins.Role) < 1 {
		log.Println("ins.Role", ins.UserName)
		return errors.New("field Role is invalid")
	}
	if len(ins.CenterName) < 1 {
		log.Println("ins.CenterName", ins.UserName)
		return errors.New("field CenterName is invalid")
	}

	return nil
}

/* */
type pongRequest struct {
	Permit string `json:"permit"`
}

type pongResponse struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Payload pong_resp `json:"payload"`
}

type pong_resp struct {
	UserName string `json:"userName"`
}

/* */
type listStaffRequest struct {
	permit
}
type listStaffResponse struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Payload list_staff_resp `json:"payload"`
}

type list_staff_resp struct {
	TotalStaff int          `json:"totalStaff"`
	ListStaff  []staff_data `json:"listStaff"`
}

type staff_data struct {
	UserName    string `json:"userName"`
	FullName    string `json:"fullName"`
	Address     string `json:"address"`
	FacultyCode string `json:"facultyCode"`
}

func withStaffModel(sm *mssql.StaffModel) staff_data {
	return staff_data{
		UserName:    sm.MaGV,
		FullName:    sm.HoTen,
		Address:     sm.DiaChi,
		FacultyCode: sm.MaKhoa,
	}
}

/* */
type listFacultyRequest struct {
	permit
}
type listFacultyResponse struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Payload list_faculty_resp `json:"payload"`
}

type list_faculty_resp struct {
	TotalFaculty int            `json:"totalFaculty"`
	ListFaculty  []faculty_data `json:"listFaculty"`
}

type faculty_data struct {
	FacultyCode string `json:"facultyCode"`
	FacultyName string `json:"facultyName"`
	CenterCode  string `json:"centerCode"`
}

func withFacultyModel(fm *mssql.FacultyModel) faculty_data {
	return faculty_data{
		FacultyName: fm.TenKH,
		FacultyCode: fm.MaKH,
		CenterCode:  fm.MaCS,
	}
}

// /* */
// type listClassRequest struct {
// 	Permit      string `json:"permit"`
// 	FacultyCode string `json:"facultyCode"`
// }
// type listClassResponse struct {
// 	Code    int             `json:"code"`
// 	Message string          `json:"message"`
// 	Payload list_class_resp `json:"payload"`
// }

// type list_class_resp struct {
// 	TotalClass int          `json:"totalClass"`
// 	ListClass  []class_data `json:"listClass"`
// }

// type class_data struct {
// 	ClassCode   string `json:"classCode"`
// 	ClassName   string `json:"className"`
// 	FacultyCode string `json:"facultyCode"`
// }

/* */
type createStaffRequest struct {
	permit
	UserName    string    `json:"userName"` // StaffCode
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	Address     string    `json:"address"`
	ClassCode   string    `json:"classCode"`
}

type createStaffResponse struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Payload create_staff `json:"payload"`
}
type create_staff struct {
}

/* */

type createCenterStaffRequest struct {
	permit
	UserName    string    `json:"userName"` // StaffCode
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	Address     string    `json:"address"`
	ClassCode   string    `json:"classCode"`
}

type createCenterStaffResponse struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Payload create_center_staff `json:"payload"`
}
type create_center_staff struct {
}
