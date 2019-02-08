

package main

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	DatabaseUser     = "postgres"
	DatabasePassword = "root"
	DatabaseName     = "quest"
)

type Db struct {
	*sql.DB
}

func newOpen() (Db, error) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmde=disable", DatabaseUser, DatabasePassword, DatabaseName)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal("Invalid DBconfig:", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("DB nreachable:", err)
	}

	return Db{db}, err
}

func main() {

	//server.Server()
	fmt.Println("ok")

}




