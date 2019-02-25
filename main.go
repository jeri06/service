

package main

import (
	"database/sql"
	"fmt"
	"github.com/jeri06/service/server"
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

	 server.Server()
	//fmt.Println("ok")
	//
	//select {
	//case messages <- msg:messages := make(chan string,1)
	//	//// signals := make(chan bool)
	//	//
	//	////go func(){
	//	//	messages <- "hi"
	//	////}()
	//	//msg := <- messages
	//	//fmt.Println(msg)
	//	fmt.Println("sent message", msg)
	//default:
	//	fmt.Println("no message sent")
	//}
}




