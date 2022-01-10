package controllers

//POST 用のfunction

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func POST_item(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/items POST")
	reqBody, _ := io.ReadAll(r.Body) // go.modに記述されているGoのバージョンを確認

	var item ItemParams
	if err := json.Unmarshal(reqBody, &item); err != nil {
		log.Fatal(err)
	}
	POST_item_for_SQL(item) //SQL操作
	json.NewEncoder(w).Encode(item)
}

func POST_item_for_SQL(item ItemParams) {
	// DataBase接続
	fmt.Println("Connect MySQL")
	db, err := sql.Open("mysql", "backend:docker@tcp(mysql_container:3306)/react_go_app")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success Connect")

	// Query
	stmt, err := db.Prepare("INSERT INTO test_table (ItemName, Price, Stock) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec(item.ItemName, item.Price, item.Stock)
}
