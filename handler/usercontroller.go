package controller

import (
	"encoding/json"
	"fmt"
	"github.com/jeri06/service/repository"
	"github.com/jeri06/service/util"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strconv"
)

func GetUser(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	uid := params["id"]
	i, err := strconv.Atoi(uid)
	if err != nil {
		fmt.Println(err)
	}
	id := uint64(i)
	data, err := repository.GetOne(id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
	b, err := json.Marshal(data)
	if err != nil {
		fmt.Println("error", err)
	}
	os.Stdout.Write(b)
	resp := util.Message(true, "success")
	resp["data"] = data
	util.Respond(w, resp)

}
