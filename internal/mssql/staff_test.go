package mssql

import (
	"context"
	"log"
	"testing"
)

func Test_GetList(t *testing.T) {
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
