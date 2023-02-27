package mssql

import (
	"context"
	wutil "csdlpt/internal/wUtil"
	"csdlpt/mssql"
	"database/sql"
	"log"
)

type loginInfo struct {
	maGV    string
	hoTen   string
	tenNhom string
}

type dbServer struct {
}

var DBServerDBC = &dbServer{}

func (ins *dbServer) Ping(ctx context.Context) (info *string, err error) {

	var (
		act = func(d *sql.DB) error {
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
		}
	)
	err = mssql.RunAdminQuery(act)
	return nil, err
}

func (ins *dbServer) GetListCenter(ctx context.Context) (list []string, err error) {

	var (
		act = func(d *sql.DB) error {
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
		}
	)
	err = mssql.RunAdminQuery(act)
	return
}

func (ins *dbServer) Login(ctx context.Context, db_permit DBPermitModel, login_name string) (
	maGv, hoTen, ho, ten, tenNhom string, data_exist bool, err error) {
	data_exist = true
	var (
		act = func(d *sql.DB) error {
			query := "USE TN_CSDLPT EXECUTE [dbo].[SP_DANGNHAP] @TENLOGIN ='" + login_name + "' ;"
			stmt, err := d.PrepareContext(ctx, query)
			if err != nil {
				return wutil.NewError(err)
			}
			defer stmt.Close()
			rows, err := stmt.QueryContext(ctx)
			if err != nil {
				return err
			}
			defer rows.Close()

			for rows.Next() {
				var (
					ho_ten sql.NullString
				)
				err := rows.Scan(&maGv, &ho_ten, &tenNhom)
				if err != nil {
					return wutil.NewError(err)
				}

				if ho_ten.Valid {
					hoTen = ho_ten.String
				}
				ho = "HO"
				ten = "TEN"

				//list = append(list, ten_cn)
			}
			return nil
		}
	)
	//err = mssql.RunAdminQuery(act)
	err = mssql.RunQuery(act, withDBConfigModel(&db_permit))
	if len(maGv) < 1 {
		data_exist = false
	}

	return
}
