package server

import (
	"github.com/jeri06/service/database"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
	"os"
	"os/signal"
	"github.com/jeri06/service/handler"
)
func DBInitialzerCaller()  {
	Db , err := database.NewOpen()
	if err != nil {
		panic(err.Error())
	}
	defer Db.Close()
	err=Db.Ping()
	if err !=nil{
		panic(err.Error())
		fmt.Println("Conected Database")
	}

}

func Server(){

	var wait time.Duration
	DBInitialzerCaller()
	router := mux.NewRouter()

	router.HandleFunc("/user/{id}", controller.GetUser).Methods("GET")

	//router.HandleFunc(“/users”, controller.CreateUser).Methods(“POST”)
	//router.HandleFunc(“/user/{id}”, controller.UpdateUser).Methods(“PUT”)
	//router.HandleFunc(“/user/{id}”, controller.DeleteUser).Methods(“DELETE”)

	srv := &http.Server{

		Handler:router,
		Addr:"127.0.0.1:7000",
		WriteTimeout: 15* time.Second,
		ReadTimeout: 15 * time.Second,
		IdleTimeout: 120 * time.Minute,
	}
	go func() {
		fmt.Println("run forever :)")
		if err := srv.ListenAndServe();
		err != nil{
			log.Println(err)
		}
	}()
	c := make(chan  os.Signal,1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx ,cancel :=context.WithTimeout(context.Background() ,wait)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Shutting down")
	os.Exit(0)
}