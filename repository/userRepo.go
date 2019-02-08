package repository

import (
	"fmt"
	"github.com/jeri06/service/database"
	"github.com/jeri06/service/model"
)

func GetAll() ([]model.User, error) {
	DB, err := database.NewOpen()
	rows, err := DB.Query("select user_name, first_name from _user limit $1", 3)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]model.User, 0)
	for rows.Next() {
		var userName string
		var firstName string
		user := model.User{}
		err := rows.Scan(&userName, &firstName)
		if err != nil {
			return nil, err
		}
		user.UserName = userName
		user.FirstName = firstName
		fmt.Printf("%8v | %6v \n", userName, firstName)
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	DB.Close()
	return users, nil
}

func GetOne(uid uint64) (model.User, error) {
	DB, err := database.NewOpen()
	user := model.User{}

	row := DB.QueryRow("select * from _user where uid = $1", uid)

	err = row.Scan(&user.UID, &user.UserName, &user.FirstName, &user.LastName, &user.MiddleName, &user.Password, &user.EmailAddress, &user.Picture, &user.Statusd, &user.Role)
	//err = row.Scan(&user.UID,&user.Statusd)
	if err != nil {
		return user, err
	}

	DB.Close()
	return user, nil
}

