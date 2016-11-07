package userservice

import (
	"database/sql"
	"fmt"
	"goWebApp/model"
	u "goWebApp/utils"
)

//UpdateUser to update user to db
func UpdateUser(db *sql.DB, user *model.FBUser) {
	tx, err := db.Begin()
	updateUser, err := db.Prepare("UPDATE fbuser SET likes=$1 WHERE id=$2")
	res, err := tx.Stmt(updateUser).Exec(user.Likes, user.ID)
	tx.Commit()
	u.CheckErr(err)
	msg := fmt.Sprintf("update new user %v", res)
	fmt.Print(msg)
}
