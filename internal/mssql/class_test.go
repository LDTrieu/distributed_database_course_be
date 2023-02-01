package mssql

import (
	"context"
	"log"
	"testing"
)

func Test_GetAllClass(t *testing.T) {
	ctx := context.Background()
	db_permit := DBPermitModel{
		CenterName: "CS1",
		UserName:   "htkn",
	}

	list, data_not_exist, err := ClassDBC.GetAll(ctx, db_permit)
	if err != nil {
		log.Println("ERR: ", err)
	}
	log.Fatal("list: ", list, "data_not_exist: ", data_not_exist)

}
