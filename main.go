package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("mysql", "root:root@(127.0.0.1)/testdb")
}

func main() {

	gRouter := mux.NewRouter()

	gRouter.HandleFunc("/", HomeHandler)

	http.ListenAndServe(":3000", gRouter)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bem vindo à homepage :D")
}
