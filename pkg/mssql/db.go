package mssql

import (
	"database/sql"
)

var (
	dbIns *sql.DB
)

// func initDB() (db *sql.DB, err error) {
// 	ctx := context.Background()
// 	dbURL := sercfg.Get(ctx, "mssql_db")
// }
