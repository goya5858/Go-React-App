package controllers

//DELETE 用のfunction

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func DELETE_item(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"] //string

	DELETE_item_for_SQL(id) //SQL操作
}

func DELETE_item_for_SQL(key string) {
	// DataBase接続
	fmt.Println("Connect MySQL")
	db, err := sql.Open("mysql", "backend:docker@tcp(mysql_container:3306)/react_go_app")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success Connect")

	// Query
	db.Exec("DELETE FROM test_table WHERE id=?", key)
}
