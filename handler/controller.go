package controller

import (
	"encoding/json"
	"fmt"
	"github.com/jeri06/service/repository"
	"github.com/jeri06/service/util"
	"net/http"
	"os"
)

func GetAll(w http.ResponseWriter, req *http.Request){
	data, err := repository.GetAllData()
	if err !=nil{
		fmt.Println(err)
	}
	b, err := json.Marshal(data)
	os.Stdout.Write(b)
	resp := util.Message(true,"success")
	resp["data"]=data
	util.Respond(w,resp)
}

