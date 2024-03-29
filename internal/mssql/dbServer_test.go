package mssql

import (
	"context"
	"log"
	"testing"
)

func Test_GetListCenter(t *testing.T) {
	ctx := context.Background()
	list, err := DBServerDBC.GetListCenter(ctx)
	if err != nil {
		log.Println("ERR: ", err)
	}
	log.Fatal("list: ", list)
}

// Ping
func Test_Ping(t *testing.T) {
	ctx := context.Background()
	list, err := DBServerDBC.Ping(ctx)
	if err != nil {
		log.Println("ERR: ", err)
	}
	log.Fatal("list: ", list)
}

// Login
func Test_Login(t *testing.T) {
	ctx := context.Background()
	userName := "th302_coso"
	db_permit := DBPermitModel{
		CenterName: "CS1",
		UserName:   "sa",
	}
	ma_gv, ho_ten, ho, ten, ten_nhom, data_exist, err := DBServerDBC.Login(ctx, db_permit, userName)
	if err != nil {
		log.Fatal(err)
	}
	if data_exist == false {
		log.Println("DATA_NOT_EXIST")
	}
	log.Println("ma_gv", ma_gv, "ho_ten", ho_ten, "ho", ho, "ten", ten, "ten_nhom", ten_nhom)
	log.Fatal("OKKK")
	//log.Fatal("login_info: ", login_info.hoTen)
}
