package model

import _ "github.com/lib/pq"
type Data struct {
	UID int `json:"id"`
	Name string `json:"name"`
	Datadetails []*Datadetails
}
type Datadetails struct {
	UID int `json:"id"`
	Dataid int `json:"dataid"`
	Detail string `json:"detail"`

}