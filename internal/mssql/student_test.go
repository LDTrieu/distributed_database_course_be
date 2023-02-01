package mssql

import (
	"context"
	"log"
	"testing"
)

func Test_GetListStudent(t *testing.T) {
	ctx := context.Background()
	db_permit := DBPermitModel{
		CenterName: "CS1",
		UserName:   "htkn",
	}

	list, data_not_exist, err := StudentDBC.GetAll(ctx, db_permit)
	if err != nil {
		log.Println("ERR: ", err)
	}
	if data_not_exist {
		log.Println("data_not_exist: ", data_not_exist)
	}
	log.Fatal("list: ", list)
}
