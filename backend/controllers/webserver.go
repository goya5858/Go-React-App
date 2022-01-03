package controllers

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

func StartWebServer() error {
	fmt.Println("Rest API with Mux Routers")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", rootPage)
	router.HandleFunc("/items", fetchAllItems).Methods("GET")
	router.HandleFunc("/item/{id}", fetchSingleItem).Methods("GET")

	router.HandleFunc("/item", createItem).Methods("POST")
	router.HandleFunc("/item/{id}", deleteItem).Methods("DELETE")
	router.HandleFunc("/item/{id}", updateItem).Methods("PUT")

	return http.ListenAndServe(fmt.Sprintf(":%d", 8080), router)
}

type ItemParams struct {
	Id       string `json:"id"`
	ItemName string `json:"item_name,omitempty"`
	Price    int    `json:"price,omitempty"`
	Stock    int    `json:"stock,omitempty"`
}

func rootPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Go Api Server")
	fmt.Println("Root endpoint is hooked!")
}

func fetchAllItems_from_SQL() []*ItemParams {
	fmt.Println("Connect MySQL")
	db, err := sql.Open("mysql", "backend:docker@tcp(mysql_container:3306)/react_go_app")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success Connect")

	rows, err := db.Query("SELECT * FROM test_table")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var items []*ItemParams
	for rows.Next() {
		var one_item ItemParams
		rows.Scan(&one_item.Id, &one_item.ItemName, &one_item.Price, &one_item.Stock)
		items = append(items, &one_item)
		fmt.Println("ID:", one_item.Id, ", ItemName:", one_item.ItemName, ", Price:", one_item.Price, ", Stock:", one_item.Stock)
	}

	return items
}

func fetchAllItems(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/items endpoint is hooked!")
	w.Header().Set("Content-Type", "application/json")

	var items []*ItemParams = fetchAllItems_from_SQL()
	json.NewEncoder(w).Encode(items)
}

func fetchSingleItem_from_SQL(key string) ItemParams {
	fmt.Println("Connect MySQL")
	db, err := sql.Open("mysql", "backend:docker@tcp(mysql_container:3306)/react_go_app")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success Connect")

	row := db.QueryRow("SELECT * FROM test_table WHERE id=?", key)
	if err != nil {
		panic(err.Error())
	}

	var one_item ItemParams
	row.Scan(&one_item.Id, &one_item.ItemName, &one_item.Price, &one_item.Stock)
	fmt.Println("ID:", one_item.Id, ", ItemName:", one_item.ItemName, ", Price:", one_item.Price, ", Stock:", one_item.Stock)

	return one_item
}

func fetchSingleItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/item endpoint is hooked!")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]

	// 現在所有してるItemsの中からIDが一致しているものを返す
	var item ItemParams = fetchSingleItem_from_SQL(key)
	json.NewEncoder(w).Encode(item)
}

func createItem_for_SQL(item ItemParams) {
	fmt.Println("Connect MySQL")
	db, err := sql.Open("mysql", "backend:docker@tcp(mysql_container:3306)/react_go_app")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success Connect")

	stmt, err := db.Prepare("INSERT INTO test_table (ItemName, Price, Stock) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec(item.ItemName, item.Price, item.Stock)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/items POST")
	reqBody, _ := io.ReadAll(r.Body) // go.modに記述されているGoのバージョンを確認

	var item ItemParams
	if err := json.Unmarshal(reqBody, &item); err != nil {
		log.Fatal(err)
	}
	createItem_for_SQL(item)
	json.NewEncoder(w).Encode(item)
}

func deleteItem_for_SQL(key string) {
	fmt.Println("Connect MySQL")
	db, err := sql.Open("mysql", "backend:docker@tcp(mysql_container:3306)/react_go_app")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success Connect")

	db.Exec("DELETE FROM test_table WHERE id=?", key)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	deleteItem_for_SQL(id)
}

func updateItem_for_SQL(key string, item ItemParams) {
	fmt.Println("Connect MySQL")
	db, err := sql.Open("mysql", "backend:docker@tcp(mysql_container:3306)/react_go_app")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success Connect")

	stmt, err := db.Prepare("UPDATE test_table SET ItemName=?, Price=?, Stock=? WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec(item.ItemName, item.Price, item.Stock, key)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	reqBody, _ := io.ReadAll(r.Body)
	var updateItem ItemParams
	if err := json.Unmarshal(reqBody, &updateItem); err != nil {
		log.Fatal(err)
	}
	updateItem_for_SQL(id, updateItem)
}
