package mssql

import (
	"context"
	"log"
	"testing"
)

func Test_GetListStaff(t *testing.T) {
	ctx := context.Background()
	db_permit := DBPermitModel{
		CenterName: "CS2",
		UserName:   "htkn",
	}

	list, err := StaffDBC.GetAll(ctx, db_permit)
	if err != nil {
		log.Println("ERR: ", err)
	}
	log.Fatal("list: ", list)
}

func Test_GetStaff(t *testing.T) {
	ctx := context.Background()
	db_permit := DBPermitModel{
		CenterName: "CS2",
		UserName:   "htkn",
	}
	// TH657
	ma_gv := "TH657"
	staff, _, err := StaffDBC.Get(ctx, db_permit, ma_gv)
	if err != nil {
		log.Println("ERR: ", err)
	}
	log.Fatal("staff: ", staff)

}
