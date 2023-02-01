package api

import (
	"csdlpt/internal/mssql"
	"errors"
	"fmt"
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

func withRequestPermission(request *loginRequest) mssql.DBPermitModel {
	return mssql.DBPermitModel{
		UserName:   "sa",
		CenterName: request.CenterName,
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
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	FullName    string `json:"fullName"`
	Address     string `json:"address"`
	FacultyCode string `json:"facultyCode"`
	StaffRole   string `json:"staffRole"`
}

func withStaffModel(sm *mssql.StaffModel) staff_data {
	return staff_data{
		UserName: sm.MaGV,
		//FullName:    fmt.Sprintf("%s%s", sm.Ho, sm.Ten),
		FullName:    sm.HoTen,
		FirstName:   sm.Ho,
		LastName:    sm.Ten,
		Address:     sm.DiaChi,
		FacultyCode: sm.MaKhoa,
	}
}

func withStaffData(sd *staff_data) mssql.StaffModel {
	return mssql.StaffModel{
		MaGV:    sd.UserName,
		Ho:      sd.FirstName,
		Ten:     sd.LastName,
		HoTen:   sd.FullName,
		DiaChi:  sd.Address,
		MaKhoa:  sd.FacultyCode,
		TenNhom: sd.StaffRole,
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
func withFacultyData(fd *faculty_data) mssql.FacultyModel {
	return mssql.FacultyModel{
		MaCS:  fd.CenterCode,
		TenKH: fd.FacultyName,
		MaKH:  fd.FacultyCode,
	}
}

/* */
type listClassRequest struct {
	permit
}
type listClassResponse struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Payload list_class_resp `json:"payload"`
}

type list_class_resp struct {
	TotalClass int          `json:"totalClass"`
	ListClass  []class_data `json:"listClass"`
}

type class_data struct {
	ClassCode   string `json:"classCode"`
	ClassName   string `json:"className"`
	FacultyCode string `json:"facultyCode"`
}

func withClassModel(cm *mssql.ClassModel) class_data {
	return class_data{
		ClassCode:   cm.MaLop,
		ClassName:   cm.TenLop,
		FacultyCode: cm.MaKH,
	}
}

func withClassData(cd *class_data) mssql.ClassModel {
	return mssql.ClassModel{
		MaLop:  cd.ClassCode,
		TenLop: cd.ClassName,
		MaKH:   cd.FacultyCode,
	}
}

/* */
type createClassRequest struct {
	permit
	ClassCode   string `json:"classCode"`
	ClassName   string `json:"className"`
	FacultyCode string `json:"facultyCode"`
}

type createClassResponse struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Payload create_class `json:"payload"`
}

type create_class struct {
}

/* */
type createStaffRequest struct {
	permit
	LoginName   string    `json:"loginName"` // StaffCode
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	Address     string    `json:"address"`
	StaffRole   string    `json:"staffRole"`
	FacultyCode string    `json:"facultyCode"`
}

type createStaffResponse struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Payload create_staff `json:"payload"`
}
type create_staff struct {
}

func (ins *createStaffRequest) validate() error {
	if len(ins.LoginName) < 1 {
		log.Println("ins.LoginName", ins.LoginName)
		return errors.New("field LoginName is invalid")
	}
	if len(ins.FirstName) < 1 {
		log.Println("ins.FirstName", ins.FirstName)
		return errors.New("field FirstName is invalid")
	}
	if len(ins.LastName) < 1 {
		log.Println("ins.LastName", ins.LastName)
		return errors.New("field LastName is invalid")
	}
	if len(ins.StaffRole) < 1 {
		log.Println("ins.StaffRole", ins.StaffRole)
		return errors.New("field StaffRole is invalid")
	}
	if len(ins.FacultyCode) < 1 {
		log.Println("ins.FacultyCode", ins.FacultyCode)
		return errors.New("field FacultyCode is invalid")
	}

	return nil
}

/* */
type createFacultyRequest struct {
	permit
	FacultyCode string `json:"facultyCode"`
	FacultyName string `json:"facultyName"`
	CenterCode  string `json:"centerCode"`
}

type createFacultyResponse struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Payload create_faculty `json:"payload"`
}

type create_faculty struct {
}

/* */
type listCourseRequest struct {
	permit
}
type listCourseResponse struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Payload list_course_resp `json:"payload"`
}

type list_course_resp struct {
	TotalCourse int           `json:"totalCourse"`
	ListCourse  []course_data `json:"listCourse"`
}

type course_data struct {
	CourseCode string `json:"courseCode"`
	CourseName string `json:"courseName"`
}

func withCourseModel(cm *mssql.CourseModel) course_data {
	return course_data{
		CourseCode: cm.MaMH,
		CourseName: cm.TenMH,
	}
}

func withCourseData(cd *course_data) mssql.CourseModel {
	return mssql.CourseModel{
		MaMH:  cd.CourseCode,
		TenMH: cd.CourseName,
	}
}

/* */
type listStudentRequest struct {
	permit
}
type listStudentResponse struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Payload list_student_resp `json:"payload"`
}

type list_student_resp struct {
	TotalStudent int            `json:"totalStudent"`
	ListStudent  []student_data `json:"listStudent"`
}

type student_data struct {
	StudentCode string    `json:"studentCode"` // userName
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	FullName    string    `json:"fullName"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	Address     string    `json:"address"`
	ClassCode   string    `json:"classCode"`
}

func withStudentModel(sm *mssql.StudentModel) student_data {
	return student_data{
		StudentCode: sm.MaSV,
		FirstName:   sm.Ho,
		LastName:    sm.Ten,
		FullName:    fmt.Sprintf("%s%s", sm.Ho, sm.Ten),
		DateOfBirth: sm.NgaySinh,
		Address:     sm.DiaChi,
		ClassCode:   sm.MaLop,
	}
}

// func withStudentData(cd *student_data) mssql.StudentModel {
// 	return mssql.StudentModel{
// 		MaMH:  cd.CourseCode,
// 		TenMH: cd.CourseName,
// 	}
// }
