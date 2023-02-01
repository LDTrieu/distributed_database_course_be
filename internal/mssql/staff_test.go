package mssql

import (
	"context"
	"fmt"
	"log"
	"strings"
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

// GetAllUserName
func Test_GetAllStaffWithoutUserName(t *testing.T) {
	ctx := context.Background()
	db_permit := DBPermitModel{
		CenterName: "CS1",
		UserName:   "htkn",
	}

	list, err := StaffDBC.GetAllStaffWithoutUserName(ctx, db_permit)
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

// CheckUserName
func Test_CheckUserName(t *testing.T) {
	ctx := context.Background()
	db_permit := DBPermitModel{
		CenterName: "CS1",
		UserName:   "htkn",
	}
	// TH204
	ma_gv := "TH219"
	staff, data_not_exist, err := StaffDBC.CheckUserName(ctx, db_permit, ma_gv)
	if err != nil {
		log.Println("ERR: ", err)
	}
	log.Println("data_not_exist", data_not_exist, "ERR: ", err)
	log.Fatal("staff: ", staff)
}

// GetStaffWithoutUserName
func Test_GetStaffWithoutUserName(t *testing.T) {
	ctx := context.Background()
	db_permit := DBPermitModel{
		CenterName: "CS1",
		UserName:   "htkn",
	}
	// TH204
	ma_gv := "TH202"
	staff, data_not_exist, err := StaffDBC.GetStaffWithoutUserName(ctx, db_permit, ma_gv)
	if err != nil {
		log.Println("ERR: ", err)
	}
	log.Println("data_not_exist", data_not_exist, "ERR: ", err)
	log.Fatal("staff: ", staff)
}

// Create
func Test_Create(t *testing.T) {
	ctx := context.Background()
	db_permit := DBPermitModel{
		CenterName: "CS1",
		UserName:   "sa",
	}
	// TH204
	staff := StaffModel{
		MaGV:   "TH209",
		Ho:     "Nguyen",
		Ten:    "Minh",
		DiaChi: "97 Man Thien",
		MaKhoa: "CNTT",
	}

	err := StaffDBC.Create(ctx, db_permit, staff)
	if err != nil {
		log.Println("ERR: ", err)
	}

	log.Fatal("OK")
}

// Test_ToLower_string
func Test_string(t *testing.T) {
	login_name := fmt.Sprintf("%s%s%s", strings.ToLower("staff.MaGV"), "_", strings.ToLower("staff.TenNhom"))
	log.Fatal(login_name)
}
