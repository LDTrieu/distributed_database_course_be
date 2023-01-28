package mssql

import (
	"context"
	"log"
	"testing"
)

func Test_GetListStaff(t *testing.T) {
	ctx := context.Background()
	db_permit := DBPermitModel{
		CenterName: "CS1",
		UserName:   "htkn",
	}

	list, err := StaffDBC.GetAll(ctx, db_permit)
	if err != nil {
		log.Println("ERR: ", err)
	}
	log.Fatal("list: ", list)
}

// GetAllUserName
func Test_GetAllUserName(t *testing.T) {
	ctx := context.Background()
	db_permit := DBPermitModel{
		CenterName: "CS1",
		UserName:   "htkn",
	}

	list, err := StaffDBC.GetAllUserName(ctx, db_permit)
	if err != nil {
		log.Println("ERR: ", err)
	}
	log.Fatal("list: ", list)
}

func Test_GetStaff(t *testing.T) {
	ctx := context.Background()
	db_permit := DBPermitModel{
		CenterName: "CS1",
		UserName:   "htkn",
	}
	// TH204
	ma_gv := "TH202"
	staff, data_not_exist, err := StaffDBC.Get(ctx, db_permit, ma_gv)
	if err != nil {
		log.Println("ERR: ", err)
	}
	log.Println("data_not_exist", data_not_exist, "ERR: ", err)
	log.Fatal("staff: ", staff)

}

// GetUserName
func Test_GetUserName(t *testing.T) {
	ctx := context.Background()
	db_permit := DBPermitModel{
		CenterName: "CS1",
		UserName:   "htkn",
	}
	// TH204
	ma_gv := "TH401"
	staff, data_not_exist, err := StaffDBC.GetUserName(ctx, db_permit, ma_gv)
	if err != nil {
		log.Println("ERR: ", err)
	}
	log.Println("data_not_exist", data_not_exist, "ERR: ", err)
	log.Fatal("staff: ", staff)

}
