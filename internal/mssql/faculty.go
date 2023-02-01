package mssql

import (
	"context"
	wutil "csdlpt/internal/wUtil"
	"csdlpt/mssql"
	"database/sql"
	"log"
)

type FacultyModel struct {
	MaKH  string `json:"maKH"`
	TenKH string `json:"tenKH"`
	MaCS  string `json:"maCS"`
}

type faculty struct {
	maKH  string
	tenKH string
	maCS  string
}

var FacultyDBC = &faculty{}

func (ins *faculty) GetAll(ctx context.Context, db_permit DBPermitModel) (list []FacultyModel, err error) {
	var (
		act = func(d *sql.DB) error {
			query := "USE TN_CSDLPT SELECT MAKH, TENKH, MACS  FROM [dbo].[KHOA];"
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
				var (
					faculty_data FacultyModel
				)
				if err := rows.Scan(&faculty_data.MaKH, &faculty_data.TenKH, &faculty_data.MaCS); err != nil {
					return wutil.NewError(err)
				}
				list = append(list, faculty_data)
			}
			return nil
		}
	)
	err = mssql.RunQuery(act, withDBConfigModel(&db_permit))
	return
}

// Get in local side
func (ins *faculty) Get(ctx context.Context, db_permit DBPermitModel, ma_kh string) (faculty FacultyModel, data_exist bool, err error) {
	var (
		act = func(d *sql.DB) error {
			query := "USE TN_CSDLPT SELECT MAKH, TENKH, MACS FROM [V_DS_KHOA]  WHERE MAKH ='DT'"
			stmt, err := d.PrepareContext(ctx, query)
			if err != nil {
				return wutil.NewError(err)
			}
			defer stmt.Close()
			rows, err := stmt.QueryContext(ctx)
			if err != nil {
				return err
			}
			defer stmt.Close()
			for rows.Next() {
				err = rows.Scan(&faculty.MaKH, &faculty.TenKH, &faculty.MaCS)
				if err != nil {
					if err == sql.ErrNoRows {
						return err
					}
					return err
				}
			}
			return nil
		}
	)
	err = mssql.RunQuery(act, withDBConfigModel(&db_permit))
	return
}

// Create
func (ins *faculty) Create(ctx context.Context, db_permit DBPermitModel, faculty FacultyModel) (err error) {

	var (
		act = func(d *sql.DB) error {
			query := " USE [TN_CSDLPT] INSERT INTO [dbo].[KHOA] ([MAKH],[TENKH],[MACS]) VALUES ('" + faculty.MaKH + "','" +
				faculty.TenKH + "','" +
				faculty.MaCS + "')"
			log.Println(query)
			_, err := d.Exec(query)
			if err != nil {
				return wutil.NewError(err)
			}

			return nil
		}
	)
	err = mssql.RunQuery(act, withDBConfigModel(&db_permit))
	return
}
