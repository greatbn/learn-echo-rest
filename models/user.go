package models

import (
	"database/sql"
	_ "database/sql"
	"fmt"

	"github.com/greatbn/learn-echo-rest/db"
)

type (
	User struct {
		Name     string `json:"name" validate:"nonzero"`
		Password string `json:"password" validate:"nonzero"`
		Email    string `json:"email" validate:"nonzero"`
	}

	Users struct {
		Users []User `json:user`
	}
)

var conn *sql.DB

func GetUser() (Users, error) {
	conn := db.CreateConn()

	sqlStatement := "select name, email from users"

	rows, err := conn.Query(sqlStatement)
	fmt.Println(rows)

	result := Users{}

	if err != nil {
		fmt.Println(err)
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.Name, &user.Email)
		if err != nil {
			fmt.Println(err)
		}
		result.Users = append(result.Users, user)
	}

	return result, nil

}
