package controllers

//GET 用のfunction ALLItems

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func GET_all_items(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/items endpoint is hooked!")
	w.Header().Set("Content-Type", "application/json")

	var items []*ItemParams = GET_all_items_from_SQL()
	json.NewEncoder(w).Encode(items)
}

func GET_all_items_from_SQL() []*ItemParams {
	// DataBase接続
	fmt.Println("Connect MySQL")
	db, err := sql.Open("mysql", "backend:docker@tcp(mysql_container:3306)/react_go_app")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success Connect")

	// Query
	rows, err := db.Query("SELECT * FROM test_table")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// 返値用のデータ作成
	var items []*ItemParams
	for rows.Next() {
		var one_item ItemParams
		rows.Scan(&one_item.Id, &one_item.ItemName, &one_item.Price, &one_item.Stock)
		items = append(items, &one_item)
		fmt.Println("ID:", one_item.Id, ", ItemName:", one_item.ItemName, ", Price:", one_item.Price, ", Stock:", one_item.Stock)
	}

	return items
}
