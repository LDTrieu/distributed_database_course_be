package mssql

import "csdlpt/mssql"

type DBPermitModel struct {
	CenterName string
	UserName   string
	PrivKey    string
}

func withDBConfigModel(dp *DBPermitModel) *mssql.DBConfigModel {
	return &mssql.DBConfigModel{
		LoginName:  dp.UserName,
		ServerName: dp.CenterName,
		PrivKey:    dp.PrivKey,
	}
}
