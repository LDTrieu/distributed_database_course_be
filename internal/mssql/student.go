package mssql

import (
	"context"
	wutil "csdlpt/internal/wUtil"
	"csdlpt/mssql"
	"database/sql"
	"time"
)

type StudentModel struct {
	MaSV     string    `json:"maSV"`
	Ho       string    `json:"ho"`
	Ten      string    `json:"ten"`
	NgaySinh time.Time `json:"ngaySinh"`
	DiaChi   string    `json:"diaChi"`
	MaLop    string    `json:"maLop"`
}

type student struct {
	maSV   string
	ho     string
	ten    string
	diaChi string
	maLop  string
}

var StudentDBC = &student{}

func (ins *student) GetAll(ctx context.Context, db_permit DBPermitModel) (list []StudentModel, data_not_exist bool, err error) {
	var (
		act = func(d *sql.DB) error {
			query := "USE TN_CSDLPT SELECT TOP (1000) [MASV], [HO], [TEN], [NGAYSINH], [DIACHI], [MALOP] FROM [TN_CSDLPT].[dbo].[SINHVIEN];"
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
					student_data StudentModel
				)
				if err := rows.Scan(&student_data.MaSV, &student_data.Ho, &student_data.Ten,
					&student_data.NgaySinh, &student_data.DiaChi, &student_data.MaLop); err != nil {
					return wutil.NewError(err)
				}
				list = append(list, student_data)
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
