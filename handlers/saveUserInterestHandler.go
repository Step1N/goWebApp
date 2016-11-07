package handlers

import (
	"database/sql"
	"encoding/json"
	"goWebApp/model"
	"goWebApp/userservice"
	u "goWebApp/utils"
	"log"
	"net/http"
)

//SaveInterestHandler to save user interest
func SaveInterestHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var fbUser *model.FBUser
		err := json.NewDecoder(r.Body).Decode(&fbUser)
		userservice.SaveUser(db, fbUser)
		u.CheckErr(err)
		err = json.NewEncoder(w).Encode(fbUser)
		u.CheckErr(err)
		log.Print("save user interest")
	}
}
