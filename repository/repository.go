package repository

import (
	"github.com/jeri06/service/database"
	"github.com/jeri06/service/model"
)

func GetAllData() ([]*model.Data, error) {
	DB, err := database.NewOpen()
	if err != nil{
		panic(err)
	}
	sql :="select * from data"
	ps :=make(map[int]*model.Data)
	rows, err := DB.Query(sql)
	if err != nil{
		panic(err)
	}

	for rows.Next(){
		p:=&model.Data{}
		rows.Scan(&p.UID,&p.Name)
		ps[p.UID]=p
	}
	sql = "select * from datadetails"
	rows, err = DB.Query(sql)
	if err != nil{
		panic(err)
	}
	for rows.Next(){
		b:=&model.Datadetails{}
		rows.Scan(&b.UID,&b.Dataid,&b.Detail)
		ps[b.Dataid].Datadetails = append(ps[b.Dataid].Datadetails, b)

	}
	datas :=make([]*model.Data,0,len(ps))

	for _, p := range ps{
		datas = append(datas,p)
	}
	return datas, err

}

//func Post()([]model.Data, error){
//
//}