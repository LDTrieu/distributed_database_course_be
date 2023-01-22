package mssql

import (
	"context"
	wutil "csdlpt/internal/wUtil"
	"csdlpt/mssql"
	"database/sql"
)

type CenterStaffModel struct {
	MaNV   string `json:"maNV"`
	HoTen  string `json:"hoTen"`
	MaCoSo string `json:"maCoSO"`
}

type centerStaff struct {
	maNV   string
	hoTen  string
	maCoSo string
}

var CenterStaffDBC = &centerStaff{}

// func (ins *centerStaff) GetAll(ctx context.Context, db_permit DBPermitModel) (list []CenterStaffModel, err error) {
// 	var (
// 		act = func(d *sql.DB) error {
// 			query := "USE TN_CSDLPT SELECT MAGV, HOTEN, MAKH FROM V_DS_GIANGVIEN;"
// 			stmt, err := d.PrepareContext(ctx, query)
// 			if err != nil {
// 				return wutil.NewError(err)
// 			}
// 			defer stmt.Close()
// 			rows, err := stmt.QueryContext(ctx)
// 			if err != nil {
// 				return wutil.NewError(err)
// 			}
// 			defer rows.Close()
// 			for rows.Next() {
// 				var (
// 					staff_data StaffModel
// 				)
// 				if err := rows.Scan(&staff_data.MaGV, &staff_data.HoTen, &staff_data.MaKhoa); err != nil {
// 					return wutil.NewError(err)
// 				}
// 				list = append(list, staff_data)
// 			}
// 			return nil
// 		}
// 	)
// 	err = mssql.RunQuery(act, withDBConfigModel(&db_permit))
// 	return
// }

// Get in local side
func (ins *centerStaff) Get(ctx context.Context, db_permit DBPermitModel, ma_nv string) (staff CenterStaffModel, data_exist bool, err error) {
	var (
		act = func(d *sql.DB) error {
			query := " USE TN_CSDLPT SELECT MANV, HOTEN, MACOSO FROM V_DS_GIANGVIEN WHERE MANV = 'TH657'"
			stmt, err := d.PrepareContext(ctx, query)
			if err != nil {
				return wutil.NewError(err)
			}
			defer stmt.Close()
			rows := stmt.QueryRowContext(ctx)
			err = rows.Scan(&staff.MaGV, &staff.HoTen, &staff.MaKhoa)
			if err != nil {
				if err == sql.ErrNoRows {
					return err
				}
				return wutil.NewError(err)
			}
			return nil
		}
	)
	err = mssql.RunQuery(act, withDBConfigModel(&db_permit))
	return
}
