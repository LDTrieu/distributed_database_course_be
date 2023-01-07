package api

import (
	"context"
	"csdlpt/internal/mssql"
	wutil "csdlpt/internal/wUtil"
	"database/sql"
	"log"
)

func __pingDB(ctx context.Context) (*pingDBResponse, error) {
	// get count(*) table
	a, err := pingHandle(ctx)
	if err != nil {
		log.Println("ERR: ", err)
	}
	log.Println(a)

	// get list table
	return nil, nil

}

func pingHandle(ctx context.Context) (info *string, err error) {
	err = mssql.RunQuery(func(d *sql.DB) error {
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
			d := ping_db_resp{}
			if err := rows.Scan(&d.DBName); err != nil {
				return wutil.NewError(err)
			}
			log.Println("OUT: ", d.DBName)
		}
		return nil
	})
	return nil, err
}

/* */
