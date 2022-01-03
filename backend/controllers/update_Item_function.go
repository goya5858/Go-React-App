package controllers

//PUT用のfunction

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func updateItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	reqBody, _ := io.ReadAll(r.Body)
	var updateItem ItemParams
	if err := json.Unmarshal(reqBody, &updateItem); err != nil {
		log.Fatal(err)
	}
	updateItem_for_SQL(id, updateItem) //SQL操作
}

func updateItem_for_SQL(key string, item ItemParams) {
	// DataBase接続
	fmt.Println("Connect MySQL")
	db, err := sql.Open("mysql", "backend:docker@tcp(mysql_container:3306)/react_go_app")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success Connect")

	// Query
	stmt, err := db.Prepare("UPDATE test_table SET ItemName=?, Price=?, Stock=? WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec(item.ItemName, item.Price, item.Stock, key)
}
