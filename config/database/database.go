package database

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"os"
	"time"
)

func DBConnection() *sql.DB {
	dbDriver := os.Getenv("DB_DRIVER")
	config := mysql.Config{
		User:                 os.Getenv("DB_USERNAME"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DB_HOSTNAME"),
		DBName:               os.Getenv("DB_NAME"),
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	DB, err := sql.Open(dbDriver, config.FormatDSN())
	if err != nil {
		panic(err)
	}
	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(100)
	DB.SetConnMaxIdleTime(5 * time.Second)
	DB.SetConnMaxLifetime(60 * time.Second)
	return DB
}
