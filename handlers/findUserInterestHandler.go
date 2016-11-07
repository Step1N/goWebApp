package handlers

import (
	"database/sql"
	"encoding/json"
	"goWebApp/model"
	s "goWebApp/userservice"
	u "goWebApp/utils"
	"log"
	"net/http"
)

//UserResponse response for user
type UserResponse struct {
	Page  int            `json:"page"`
	Users []model.FBUser `json:"users"`
}

//FindInterestHandler to find user interest
func FindInterestHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users := s.ReadUser(db)
		err := json.NewEncoder(w).Encode(users)
		u.CheckErr(err)
		log.Print("Writen on response")
	}
}
