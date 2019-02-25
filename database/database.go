package database

import (
	"database/sql"
	"fmt"

	"log"
)

const (
	DatabaseUser     = "postgres"
	DatabasePassword = "root"
	DatabaseName     = "golang"
)

type Db struct {
	*sql.DB
}

func NewOpen() (Db,error) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DatabaseUser, DatabasePassword, DatabaseName)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal("Invalid DBconfig:", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("DB nreachable:", err)
	}
	//dbinfo :=pg.Connect(&pg.Options{
	//	User:     "postgres",
	//	Password: "root",
	//	Database: "golang",
	//})
	//defer dbinfo.Close()


	return Db{db}, err
}
