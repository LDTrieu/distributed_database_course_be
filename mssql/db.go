package mssql

import (
	wutil "csdlpt/internal/wUtil"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/spf13/viper"
)

var (
	dbIns, readOnlyDbIns *sql.DB
)

// func initApp() {
// 	cfgBuff, err := ioutil.ReadFile("cfg.yaml")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	appCfg = &cfg{}
// 	if err := yaml.Unmarshal(cfgBuff, appCfg); err != nil {
// 		log.Fatal(err)
// 	}
// }

// func main() {
// 	initApp()

//		connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d", *server, *user, *password, *port)
//		conn, err := sql.Open("mssql", connString)
//		if err != nil {
//			log.Fatal("Open connection failed:", err.Error())
//		}
//		defer conn.Close()
//		stmt, err := conn.Prepare("SELECT COUNT(name) FROM master.sys.databases;")
//		if err != nil {
//			log.Fatal("Prepare failed:", err.Error())
//		}
//		defer stmt.Close()
//		row := stmt.QueryRow()
//		var somenumber string
//		var somechars string
//		err = row.Scan(&somechars)
//		if err != nil {
//			log.Fatal("Scan failed:", err.Error())
//		}
//		fmt.Printf("somenumber:%s\n", somenumber)
//		fmt.Printf("somechars:%s\n", somechars)
//		// router := gin.Default()
//		log.Println("RUN MAIN")
//	}
// func initApp() {
// 	viper.SetConfigName("config")
// 	viper.AddConfigPath(".")
// 	if err := viper.ReadInConfig(); err != nil {
// 		log.Fatalf("Error while reading config file %s", err)
// 	}
// 	server := viper.Sub("mssql").GetString("server")
// 	user := viper.Sub("mssql").GetString("user")
// 	password := viper.Sub("mssql").GetString("password")
// 	port := viper.Sub("mssql").GetString("port")

// 	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s", server, user, password, port)
// 	log.Println("connString", connString)
// 	db, err := sql.Open("mssql", connString)
// 	if err != nil {
// 		log.Fatal("Open connection failed:", err.Error())
// 	}
// 	db.SetConnMaxLifetime(time.Minute * 3)
// 	db.SetMaxOpenConns(10)
// 	db.SetMaxIdleConns(5)
// 	if err = db.Ping(); err != nil {
// 		err = wutil.NewError(err)
// 	}
// 	return
// }

func initDB(db_config *DBConfigModel) (db *sql.DB, err error) {

	// ctx := context.Background()
	var (
		server string
	)
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	//user := viper.Sub("mssql").GetString("user")
	user := db_config.LoginName
	password := viper.Sub("mssql").GetString("password")
	port := viper.Sub("mssql").GetString("port")
	//		conn, err := sql.Open("mssql", connString)
	switch db_config.ServerName {
	case "CS1":
		log.Println("CS1")
		server = viper.Sub("mssql").GetString("server1")
	case "CS2":
		log.Println("CS2")
		server = viper.Sub("mssql").GetString("server2")
	default:
		server = viper.Sub("mssql").GetString("server")
	}
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s", server, user, password, port)
	log.Println("connString", connString)
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	if err = db.Ping(); err != nil {
		err = wutil.NewError(err)
	}
	return
}

func initAdminDB() (db *sql.DB, err error) {

	// ctx := context.Background()
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	server := viper.Sub("mssql").GetString("server")
	user := viper.Sub("mssql").GetString("user")
	password := viper.Sub("mssql").GetString("password")
	port := viper.Sub("mssql").GetString("port")
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s", server, user, password, port)
	log.Println("connString", connString)
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	if err = db.Ping(); err != nil {
		err = wutil.NewError(err)
	}
	return
}
func runQuery(act func(*sql.DB) error, db_config *DBConfigModel) (err error) {
	if dbIns == nil {
		dbIns, err = initDB(db_config)
		if err != nil {
			return
		}
	}
	err = act(dbIns)
	return
}
func runAdminQuery(act func(*sql.DB) error) (err error) {
	if dbIns == nil {
		dbIns, err = initAdminDB()
		if err != nil {
			return
		}
	}
	err = act(dbIns)
	return
}

func RunQuery(act func(db *sql.DB) error, db_config *DBConfigModel) error {
	return runQuery(act, db_config)
}

func RunAdminQuery(act func(db *sql.DB) error) error {
	return runAdminQuery(act)
}
