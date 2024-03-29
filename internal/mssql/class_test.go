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

func Test_CreateClass(t *testing.T) {
	var (
		ctx       = context.Background()
		db_permit = DBPermitModel{
			CenterName: "CS1",
			UserName:   "th301_coso",
		}
		class_data = ClassModel{
			MaLop:  "CS4",
			TenLop: "Vient Thong T 3",
			MaKH:   "CNTT",
		}
	)

	if err := ClassDBC.Create(ctx, db_permit, class_data); err != nil {
		log.Fatal(err)
	}

}
