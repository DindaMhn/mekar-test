package config

import (
	"database/sql"
	"fmt"
	"log"
	"mekar/master/tools/getEnv"

	//for connection mysql db
	_ "github.com/go-sql-driver/mysql"
)

//EnvConn is function to get Environment Variabel for DB connection
func EnvConn() *sql.DB {
	dbEngine := getEnv.ViperGetEnv("DB_ENGINE", "mysql") //mysql
	dbUser := getEnv.ViperGetEnv("DB_USER", "root")      //root
	dbPassword := getEnv.ViperGetEnv("DB_PASSWORD", "dindarani1")
	dbHost := getEnv.ViperGetEnv("DB_HOST", "localhost") //localhost
	dbPort := getEnv.ViperGetEnv("DB_PORT", "3306")      //3306
	dbSchema := getEnv.ViperGetEnv("DB_SCHEMA", "pencatatan_keuangan")

	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbSchema)
	db, err := sql.Open(dbEngine, dbSource)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}
