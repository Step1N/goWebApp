package handlers

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"goWebApp/model"
	"goWebApp/userservice"
	u "goWebApp/utils"
	"log"
	"net/http"
	"strconv"
)

//UpdateInterestHandler to update user interest
func UpdateInterestHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var fbUser *model.FBUser
		err := json.NewDecoder(r.Body).Decode(&fbUser)
		vars := mux.Vars(r)
		id := vars["id"]
		userID, _ := strconv.Atoi(id)
		fbUser.ID = userID
		userservice.UpdateUser(db, fbUser)
		u.CheckErr(err)
		err = json.NewEncoder(w).Encode(fbUser)
		u.CheckErr(err)
		log.Printf("updated user interest %v", err)
	}
}
