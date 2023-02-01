package mssql

import (
	"context"
	wutil "csdlpt/internal/wUtil"
	"csdlpt/mssql"
	"database/sql"
)

type ClassModel struct {
	MaLop  string `json:"maLop"`
	TenLop string `json:"tenLop"`
	MaKH   string `json:"maKh"`
}

type class struct {
	maLop  string
	tenLop string
	maKh   string
}

var ClassDBC = &class{}

func (ins *class) GetAll(ctx context.Context, db_permit DBPermitModel) (list []ClassModel, data_not_exist bool, err error) {
	var (
		act = func(d *sql.DB) error {
			query := "USE TN_CSDLPT SELECT TOP (1000) [TENLOP], [TENLOP], [MAKH] FROM [LOP];"
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
					class_data ClassModel
				)
				if err := rows.Scan(&class_data.MaLop, &class_data.TenLop, &class_data.MaKH); err != nil {
					return wutil.NewError(err)
				}
				list = append(list, class_data)
			}
			return nil
		}
	)
	err = mssql.RunQuery(act, withDBConfigModel(&db_permit))
	if err == sql.ErrNoRows {
		return nil, true, nil
	}
	return
}
