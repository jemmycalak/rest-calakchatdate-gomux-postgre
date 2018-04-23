package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func GetPostgresDB() (*sql.DB, error) {
	host := "127.0.0.1"
	user := "postgres"
	password := "postgre"
	dbname := "dbgolang"

	desc := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbname)

	db, err := createConection(desc)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func createConection(desc string) (*sql.DB, error) {
	db, err := sql.Open("postgres", desc) //jenis databasenya apa,, kalau myql ganti dg mysqldriver

	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	return db, nil
}
