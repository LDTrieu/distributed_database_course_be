package api

import (
	"errors"
	"log"
)

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
