package handlers

import (
	"fmt"
	"net/http"

	"database/sql"
	gh "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//APIContext context about router
type APIContext struct {
	DB *sql.DB
}

//Context for router
func Context(ctx *APIContext) http.Handler {

	r := mux.NewRouter()
	findInterestHandler := FindInterestHandler(ctx.DB)
	saveInterestHandler := SaveInterestHandler(ctx.DB)
	updateInterestHandler := UpdateInterestHandler(ctx.DB)
	r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to Home")
	})

	r.HandleFunc("/userInterest", findInterestHandler).Methods("GET")
	r.HandleFunc("/userInterest", saveInterestHandler).Methods("POST")
	r.HandleFunc("/userInterest/{id}", updateInterestHandler).Methods("PUT")

	return gh.CompressHandler(r)
}
