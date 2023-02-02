package mssql

import (
	"context"
	"log"
	"testing"
)

func Test_GetAllCourse(t *testing.T) {
	ctx := context.Background()
	db_permit := DBPermitModel{
		CenterName: "CS1",
		UserName:   "th301_coso",
	}

	list, data_not_exist, err := CourseDBC.GetAll(ctx, db_permit)
	if err != nil {
		log.Println("ERR: ", err)
	}
	log.Fatal("list: ", list, "data_not_exist: ", data_not_exist)

}

func Test_CreateCourse(t *testing.T) {
	var (
		ctx       = context.Background()
		db_permit = DBPermitModel{
			CenterName: "CS1",
			UserName:   "th301_coso",
		}
		course_data = CourseModel{
			MaMH:  "CS4",
			TenMH: "Vient Thong T 3",
		}
	)

	if err := CourseDBC.Create(ctx, db_permit, course_data); err != nil {
		log.Fatal(err)
	}

}
