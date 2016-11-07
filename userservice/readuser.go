package userservice

import (
	"database/sql"
	"fmt"
	"goWebApp/model"
	u "goWebApp/utils"
)

//ReadUser to fetch user from db
func ReadUser(db *sql.DB) []*model.FBUser {
	rows, err := db.Query("SELECT * FROM fbuser")
	u.CheckErr(err)
	fbUsers := []*model.FBUser{}
	for rows.Next() {
		fbUser := &model.FBUser{}
		err = rows.Scan(&fbUser.ID, &fbUser.Name, &fbUser.Interest, &fbUser.Likes)
		u.CheckErr(err)
		fmt.Printf("Read all users in loop %v", fbUser)
		fbUsers = append(fbUsers, fbUser)
	}

	fmt.Printf("Read all users %v", fbUsers)
	return fbUsers
}
