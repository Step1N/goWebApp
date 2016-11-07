package main

import (
	"fmt"

	"database/sql"
	_ "github.com/lib/pq"
	"github.com/urfave/negroni"

	h "goWebApp/handlers"
	u "goWebApp/utils"
)

const (
	//DBUser User name for db
	DBUser = "admin"
	//DBPassword Password  for db
	DBPassword = "admin"
	//DBName  name for db
	DBName = "goWebApp"
)

func main() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DBUser, DBPassword, DBName)
	db, err := sql.Open("postgres", dbinfo)
	u.CheckErr(err)
	defer db.Close()

	n := negroni.New()
	apiContext := &h.APIContext{DB: db}
	routes := h.Context(apiContext)

	n.UseHandler(routes)
	n.Run(":8080")
	fmt.Println("Serving On 8080")
}
