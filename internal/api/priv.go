package api

import (
	"context"
	"csdlpt/internal/mssql"
	"csdlpt/model"
	"log"
)

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

func __login(ctx context.Context, request *loginRequest) (
	*loginResponse, error) {
	// Validate request
	if err := request.validate(); err != nil {
		return &loginResponse{
			Code:    model.StatusBadRequest,
			Message: "DATA_INVALID",
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
		Payload: login_resp{},
	}, nil
}
