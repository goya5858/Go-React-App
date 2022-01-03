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

func connectMySQL() {
	fmt.Println("Connect MySQL")
	db, err := sql.Open("mysql", "backend:docker@tcp(mysql_container:3306)/react_go_app")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success Connect")
	fmt.Println("Changed")
	fmt.Println(db)

	// rowが一つのみだとわかっているとき
	//row := db.QueryRow("SELECT * FROM test_table")
	//var one_item ItemParams
	//row.Scan(&one_item.Id, &one_item.ItemName, &one_item.Price, &one_item.Stock)
	//fmt.Println("ID:", one_item.Id, ", ItemName:", one_item.ItemName, ", Price:", one_item.Price, ", Stock:", one_item.Stock)

	rows, err := db.Query("SELECT * FROM test_table")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var one_item ItemParams
		rows.Scan(&one_item.Id, &one_item.ItemName, &one_item.Price, &one_item.Stock)
		fmt.Println("ID:", one_item.Id, ", ItemName:", one_item.ItemName, ", Price:", one_item.Price, ", Stock:", one_item.Stock)
	}
}

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

var items []*ItemParams

func rootPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Go Api Server")
	fmt.Println("Root endpoint is hooked!")

}

func fetchAllItems(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/items endpoint is hooked!")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)

	connectMySQL()
}

func fetchSingleItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/item endpoint is hooked!")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]

	// 現在所有してるItemsの中からIDが一致しているものを返す
	for _, item := range items {
		if item.Id == key {
			json.NewEncoder(w).Encode(item)
		}
	}
}

func createItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/items POST")
	reqBody, _ := io.ReadAll(r.Body) // go.modに記述されているGoのバージョンを確認
	//reqBody, _ := ioutil.ReadAll(r.Body)
	var item ItemParams
	if err := json.Unmarshal(reqBody, &item); err != nil {
		log.Fatal(err)
	}

	items = append(items, &item)
	json.NewEncoder(w).Encode(item)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, item := range items {
		if item.Id == id {
			// indexがidの部分を飛ばす
			items = append(items[:index], items[index+1:]...)
		}
	}
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	reqBody, _ := io.ReadAll(r.Body)
	//reqBody, _ := ioutil.ReadAll(r.Body)
	var updateItem ItemParams
	if err := json.Unmarshal(reqBody, &updateItem); err != nil {
		log.Fatal(err)
	}

	for index, item := range items {
		if item.Id == id {
			items[index] = &ItemParams{
				Id:       item.Id,
				ItemName: updateItem.ItemName,
				Price:    updateItem.Price,
				Stock:    updateItem.Stock,
			}
		}
	}
}

func init() {
	items = []*ItemParams{
		{
			Id:       "1",
			ItemName: "item_1",
			Price:    2500,
			Stock:    100,
		},
		{
			Id:       "2",
			ItemName: "item_2",
			Price:    1200,
			Stock:    200,
		},
	}
}
