package model

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type User struct {
	UID          uint64
	UserName     string         `json:"username"`
	Password     string         `json:"password"`
	FirstName    string         `json:"first_name"`
	LastName     string         `json:"last_name"`
	MiddleName   string         `json:"middle_name"`
	EmailAddress string         `json:"email_address"`
	Picture      sql.NullString
	Statusd      string         `json:"statusd"`
	Role         string         `json:"role"`
}

func NewUser(
	uid uint64,
	userName,
	password,
	fName,
	LName string,
	mName string,
	emailAddress string,
	picture sql.NullString,
	statusd string,
	role string) User {
	user := new(User)
	user.UID = uid
	user.UserName = userName
	user.Password = password
	user.FirstName = fName
	user.LastName = LName
	user.MiddleName = mName
	user.EmailAddress = emailAddress
	user.Picture = picture
	user.Statusd = statusd
	user.Role = role
	//user.DateJoined = dateJoined
	return *user
}
