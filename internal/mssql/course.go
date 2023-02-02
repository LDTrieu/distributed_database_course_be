package mssql

import (
	"context"
	wutil "csdlpt/internal/wUtil"
	"csdlpt/mssql"
	"database/sql"
	"log"
)

type CourseModel struct {
	MaMH  string `json:"maMH"`
	TenMH string `json:"tenMH"`
}

type course struct {
	maMH  string
	tenMH string
}

var CourseDBC = &course{}

func (ins *course) GetAll(ctx context.Context, db_permit DBPermitModel) (list []CourseModel, data_not_exist bool, err error) {
	var (
		act = func(d *sql.DB) error {
			query := "USE TN_CSDLPT SELECT TOP (1000) [MAMH], [TENMH] FROM [MONHOC];"
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
					course_data CourseModel
				)
				if err := rows.Scan(&course_data.MaMH, &course_data.TenMH); err != nil {
					return wutil.NewError(err)
				}
				list = append(list, course_data)
			}
			return nil
		}
	)
	err = mssql.RunQuery(act, withDBConfigModel(&db_permit))
	if err == sql.ErrNoRows {
		return nil, true, nil
	}
	if len(list) == 0 {
		return nil, true, err
	}
	return
}

func (ins *course) Create(ctx context.Context, db_permit DBPermitModel, course CourseModel) (err error) {

	var (
		act = func(d *sql.DB) error {
			query := "USE [TN_CSDLPT] INSERT INTO [dbo].[MONHOC] ([MAMH],[TENMH]) VALUES ('" + course.MaMH + "','" + course.TenMH + "')"
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
