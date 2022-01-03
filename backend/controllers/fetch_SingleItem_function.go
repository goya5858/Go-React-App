package controllers

//GET 用のfunction Single_item

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func fetchSingleItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/item endpoint is hooked!")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]

	var item ItemParams = fetchSingleItem_from_SQL(key) //SQL操作
	json.NewEncoder(w).Encode(item)
}

func fetchSingleItem_from_SQL(key string) ItemParams {
	// DataBase接続
	fmt.Println("Connect MySQL")
	db, err := sql.Open("mysql", "backend:docker@tcp(mysql_container:3306)/react_go_app")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success Connect")

	// Query
	row := db.QueryRow("SELECT * FROM test_table WHERE id=?", key)
	if err != nil {
		panic(err.Error())
	}

	// 返値用のデータ作成
	var one_item ItemParams
	row.Scan(&one_item.Id, &one_item.ItemName, &one_item.Price, &one_item.Stock)
	fmt.Println("ID:", one_item.Id, ", ItemName:", one_item.ItemName, ", Price:", one_item.Price, ", Stock:", one_item.Stock)

	return one_item
}
