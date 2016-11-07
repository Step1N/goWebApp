package userservice

import (
	"database/sql"
	"fmt"
	"goWebApp/model"
	u "goWebApp/utils"
)

//SaveUser to save user to db
func SaveUser(db *sql.DB, user *model.FBUser) {
	var id int
	err := db.QueryRow("INSERT INTO fbuser(name,interest,likes) VALUES($1,$2,$3) returning id;", user.Name, user.Interest, user.Likes).Scan(&id)
	u.CheckErr(err)
	msg := fmt.Sprintf("added new user %v", id)
	fmt.Print(msg)
}
