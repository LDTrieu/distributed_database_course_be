package mssql

import (
	"context"
	"log"
	"testing"
)

func Test_GetListFaculty(t *testing.T) {
	ctx := context.Background()
	db_permit := DBPermitModel{
		CenterName: "CS2",
		UserName:   "htkn",
	}

	list, err := FacultyDBC.GetAll(ctx, db_permit)
	if err != nil {
		log.Println("ERR: ", err)
	}
	log.Fatal("list: ", list)
}

func Test_GetFaculty(t *testing.T) {
	ctx := context.Background()
	db_permit := DBPermitModel{
		CenterName: "CS2",
		UserName:   "htkn",
	}
	// TH657
	ma_kh := "TH657"
	faculty, data_exist, err := FacultyDBC.Get(ctx, db_permit, ma_kh)
	if err != nil {
		log.Println("ERR: ", err)
	}
	log.Fatal("faculty: ", faculty, "data_exist: ", data_exist)
}
