package mssql

import (
	"context"
	wutil "csdlpt/internal/wUtil"
	"csdlpt/mssql"
	"database/sql"
	"log"
)

//	type dbServer struct {
//		ID int64 `json:"id"`
//	}
type loginInfo struct {
	maGV    string
	hoTen   string
	tenNhom string
}

type dbServer struct {
}

var DBServerDBC = &dbServer{}

func (ins *dbServer) Ping(ctx context.Context) (
	info *string, err error) {
	err = mssql.RunQuery(func(d *sql.DB) error {
		query := "USE TN_CSDLPT select MACS from COSO"
		stmt, err := d.PrepareContext(ctx, query)
		if err != nil {
			return wutil.NewError(err)
		}
		log.Println("stmt", stmt)
		defer stmt.Close()
		rows, err := stmt.QueryContext(ctx)
		if err != nil {
			return wutil.NewError(err)
		}
		log.Println("rows", rows)
		defer rows.Close()
		for rows.Next() {
			var d string
			if err := rows.Scan(&d); err != nil {
				return wutil.NewError(err)
			}
			log.Println("OUT: ", d)
		}
		return nil
	})
	return nil, err
}

func (ins *dbServer) GetListCenter(ctx context.Context) (
	list []string, err error) {
	err = mssql.RunQuery(func(d *sql.DB) error {
		query := "USE TN_CSDLPT SELECT TENCN FROM V_DS_PHANMANH"
		stmt, err := d.PrepareContext(ctx, query)
		if err != nil {
			return wutil.NewError(err)
		}
		defer stmt.Close()
		rows, err := stmt.QueryContext(ctx)
		if err != nil {
			return wutil.NewError(err)
		}
		defer rows.Close()
		for rows.Next() {
			var ten_cn string
			if err := rows.Scan(&ten_cn); err != nil {
				return wutil.NewError(err)
			}
			list = append(list, ten_cn)
		}
		return nil
	})

	return
}

func (ins *dbServer) Login(ctx context.Context, userName string) (
	login_info *loginInfo, err error) {
	// if len(userName) < 1 {
	// 	return wutil.NewError("err")
	// }
	err = mssql.RunQuery(func(d *sql.DB) error {
		query := "USE TN_CSDLPT EXECUTE [dbo].[SP_DANGNHAP] @TENLOGIN ='" + userName + "' ;"
		stmt, err := d.PrepareContext(ctx, query)
		if err != nil {
			return wutil.NewError(err)
		}
		log.Println("userName", userName)
		defer stmt.Close()
		rows, err := stmt.QueryContext(ctx, userName)
		if err != nil {
			log.Println("ERR: ", err)
			return err
		}
		log.Println("rows: ", rows)
		defer rows.Close()
		log.Println("LINE 96")
		for rows.Next() {
			log.Println("LINE 98")
			var (
				ma_gv, ho_ten, ten_nhom sql.NullString
			)
			err := rows.Scan(&ma_gv, &ho_ten, &ten_nhom)
			if err != nil {
				log.Println("LINE 103")
				return wutil.NewError(err)
			}
			log.Println("ma_gv: ", ma_gv.String, "ho_ten: ", ho_ten.String, "ten_nhom: ", ten_nhom.String)
			if ma_gv.Valid {
				login_info.maGV = ma_gv.String
			}
			if ho_ten.Valid {
				login_info.hoTen = ho_ten.String
			}
			if ten_nhom.Valid {
				login_info.tenNhom = ten_nhom.String
			}
			log.Println("LINE 113")
			log.Println("ma_gv: ", ma_gv.String, "ho_ten: ", ho_ten.String, "ten_nhom: ", ten_nhom.String)
			//list = append(list, ten_cn)

		}
		log.Println("END: ")
		return nil
	})
	log.Println("login_info: ", login_info.hoTen)

	return
}
