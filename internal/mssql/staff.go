package mssql

import (
	"context"
	wutil "csdlpt/internal/wUtil"
	"csdlpt/mssql"
	"database/sql"
)

type StaffModel struct {
	MaGV    string `json:"maGV"`
	HoTen   string `json:"hoTen"`
	Ho      string `json:"ho"`
	Ten     string `json:"ten"`
	TenNhom string `json:"tenNhom"`
	DiaChi  string `json:"diaChi"`
	MaKhoa  string `json:"maKhoa"`
}

type staff struct {
	maGV    string
	hoTen   string
	tenNhom string
	diaChi  string
	maKhoa  string
}

var StaffDBC = &staff{}

func (ins *staff) GetAll(ctx context.Context, db_permit DBPermitModel) (list []StaffModel, err error) {
	var (
		act = func(d *sql.DB) error {
			query := "USE TN_CSDLPT SELECT MAGV, HOTEN, MAKH FROM V_DS_GIANGVIEN;"
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
					staff_data StaffModel
				)
				if err := rows.Scan(&staff_data.MaGV, &staff_data.HoTen, &staff_data.MaKhoa); err != nil {
					return wutil.NewError(err)
				}
				list = append(list, staff_data)
			}
			return nil
		}
	)
	err = mssql.RunQuery(act, withDBConfigModel(&db_permit))
	return
}

func (ins *staff) GetAllUserName(ctx context.Context, db_permit DBPermitModel) (list []StaffModel, err error) {
	var (
		act = func(d *sql.DB) error {
			query := "USE TN_CSDLPT SELECT MAGV, HOTEN, MAKH FROM V_DS_USERNAME;"
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
					staff_data StaffModel
				)
				if err := rows.Scan(&staff_data.MaGV, &staff_data.HoTen, &staff_data.MaKhoa); err != nil {
					return wutil.NewError(err)
				}
				list = append(list, staff_data)
			}
			return nil
		}
	)
	err = mssql.RunQuery(act, withDBConfigModel(&db_permit))
	return
}

// GetAllStaffWithoutUserName
func (ins *staff) GetAllStaffWithoutUserName(ctx context.Context, db_permit DBPermitModel) (list []StaffModel, err error) {
	var (
		act = func(d *sql.DB) error {
			query := "USE TN_CSDLPT SELECT MAGV, HOTEN, MAKH FROM V_DS_GIANGVIEN_2;"
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
					staff_data StaffModel
				)
				if err := rows.Scan(&staff_data.MaGV, &staff_data.HoTen, &staff_data.MaKhoa); err != nil {
					return wutil.NewError(err)
				}
				list = append(list, staff_data)
			}
			return nil
		}
	)
	err = mssql.RunQuery(act, withDBConfigModel(&db_permit))
	return
}

// Get in local side
func (ins *staff) Get(ctx context.Context, db_permit DBPermitModel, ma_gv string) (staff StaffModel, data_not_exist bool, err error) {
	var (
		act = func(d *sql.DB) error {
			query := " USE TN_CSDLPT SELECT MAGV, HOTEN, MAKH FROM V_DS_GIANGVIEN WHERE MAGV = '" + ma_gv + "';"
			stmt, err := d.PrepareContext(ctx, query)
			if err != nil {
				return wutil.NewError(err)
			}
			defer stmt.Close()
			rows := stmt.QueryRowContext(ctx)
			err = rows.Scan(&staff.MaGV, &staff.HoTen, &staff.MaKhoa)
			if err != nil {
				return err
			}
			return nil
		}
	)
	err = mssql.RunQuery(act, withDBConfigModel(&db_permit))
	if err == sql.ErrNoRows {
		return staff, true, nil
	}
	return
}

// Get in local side
func (ins *staff) CheckUserName(ctx context.Context, db_permit DBPermitModel, ma_gv string) (staff StaffModel, data_not_exist bool, err error) {
	var (
		act = func(d *sql.DB) error {
			query := " USE TN_CSDLPT SELECT MAGV, HOTEN, MAKH FROM V_DS_USERNAME WHERE MAGV = '" + ma_gv + "';"
			stmt, err := d.PrepareContext(ctx, query)
			if err != nil {
				return wutil.NewError(err)
			}
			defer stmt.Close()
			rows := stmt.QueryRowContext(ctx)
			err = rows.Scan(&staff.MaGV, &staff.HoTen, &staff.MaKhoa)
			if err != nil {
				return err
			}
			return nil
		}
	)
	err = mssql.RunQuery(act, withDBConfigModel(&db_permit))
	if err == sql.ErrNoRows {
		return staff, true, nil
	}
	return
}

func (ins *staff) GetUserName(ctx context.Context, db_permit DBPermitModel, ma_gv string) (staff StaffModel, data_not_exist bool, err error) {
	var (
		act = func(d *sql.DB) error {
			query := " USE TN_CSDLPT SELECT MAGV, HOTEN, MAKH FROM V_DS_USERNAME WHERE MAGV = '" + ma_gv + "';"
			stmt, err := d.PrepareContext(ctx, query)
			if err != nil {
				return wutil.NewError(err)
			}
			defer stmt.Close()
			rows := stmt.QueryRowContext(ctx)
			err = rows.Scan(&staff.MaGV, &staff.HoTen, &staff.MaKhoa)
			if err != nil {
				return err
			}
			return nil
		}
	)
	err = mssql.RunQuery(act, withDBConfigModel(&db_permit))
	if err == sql.ErrNoRows {
		return staff, true, nil
	}
	return
}

func (ins *staff) GetStaffWithoutUserName(ctx context.Context, db_permit DBPermitModel, ma_gv string) (staff StaffModel, data_not_exist bool, err error) {
	var (
		act = func(d *sql.DB) error {
			query := " USE TN_CSDLPT SELECT MAGV, HOTEN, MAKH FROM V_DS_GIANGVIEN_2 WHERE MAGV = '" + ma_gv + "';"
			stmt, err := d.PrepareContext(ctx, query)
			if err != nil {
				return wutil.NewError(err)
			}
			defer stmt.Close()
			rows := stmt.QueryRowContext(ctx)
			err = rows.Scan(&staff.MaGV, &staff.HoTen, &staff.MaKhoa)
			if err != nil {
				return err
			}
			return nil
		}
	)
	err = mssql.RunQuery(act, withDBConfigModel(&db_permit))
	if err == sql.ErrNoRows {
		return staff, true, nil
	}
	return
}

// Create
func (ins *staff) Create(ctx context.Context, db_permit DBPermitModel, staff StaffModel) (err error) {

	var (
		act = func(d *sql.DB) error {
			query := " USE TN_CSDLPT EXEC SP_CREATE_GIANGVIEN @MAGV ='" +
				staff.MaGV + "',@HO ='" +
				staff.Ho + "',@TEN='" +
				staff.Ten + "',@DIACHI ='" +
				staff.DiaChi + "', @MAKH ='" +
				staff.MaKhoa + "'"
			//query := "USE TN_CSDLPT EXEC SP_CREATE_GIANGVIEN @MAGV ='TH208',@HO ='PHAM',@TEN='HANNI',@DIACHI ='Australian', @MAKH ='CNTT'"
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

// Create Login
func (ins *staff) CreateLogin(ctx context.Context, db_permit DBPermitModel, staff StaffModel) (err error) {

	var (
		act = func(d *sql.DB) error {
			query := " USE TN_CSDLPT EXEC SP_CREATE_GIANGVIEN @MAGV ='" +
				staff.MaGV + "',@HO ='" +
				staff.Ho + "',@TEN='" +
				staff.Ten + "',@DIACHI ='" +
				staff.DiaChi + "', @MAKH ='" +
				staff.MaKhoa + "'"
			//query := "USE TN_CSDLPT EXEC SP_CREATE_GIANGVIEN @MAGV ='TH208',@HO ='PHAM',@TEN='HANNI',@DIACHI ='Australian', @MAKH ='CNTT'"
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
