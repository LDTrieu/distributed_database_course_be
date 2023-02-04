package mssql

type QuestionModel struct {
	MaCH    string `json:"maCH"`
	MaMH    string `json:"maMH"`
	TrinhDo string `json:"trinhDo"`
	NoiDung string `json:"noiDung"`
	ChooseA string `json:"chooseA"`
	ChooseB string `json:"chooseB"`
	ChooseC string `json:"chooseC"`
	ChooseD string `json:"chooseD"`
}

type question struct {
	maCH    string
	maMH    string
	trinhDo string
	noiDung string
	chooseA string
	chooseB string
	chooseC string
	chooseD string
}

var QuestionDBC = &question{}

// Get in local side
// func (ins *question) GetByStaffCode(ctx context.Context, db_permit DBPermitModel, ma_gv string) (list []QuestionModel, err error) {
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
