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
	userName := "htkn"
	login_info, err := DBServerDBC.Login(ctx, userName)
	if err != nil {
		log.Println("ERR: ", err)
	}
	log.Fatal("login_info: ", login_info.hoTen)
}
