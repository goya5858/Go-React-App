package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func connectMySQL() {
	DBMS := "mysql"   // どのSQLを使用するか　今回はMySQL
	USER := "backend" // backend_containerから利用
	PASS := "docker"
	PROTOCOL := "tcp(mysql:3306)"
	DBNAME := "react_go_app"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	fmt.Println("Connect MySQL")
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success Connect")
	fmt.Println("Changed")
	fmt.Println(db)
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
	Id        string    `json:"id"`
	ItemName  string    `json:"item_name,omitempty"`
	Price     int       `json:"price,omitempty"`
	Stock     int       `json:"stock,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
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
				Id:        item.Id,
				ItemName:  updateItem.ItemName,
				Price:     updateItem.Price,
				Stock:     updateItem.Stock,
				CreatedAt: item.CreatedAt,
				UpdatedAt: updateItem.UpdatedAt,
				DeletedAt: item.DeletedAt,
			}
		}
	}
}

func init() {
	items = []*ItemParams{
		{
			Id:        "1",
			ItemName:  "item_1",
			Price:     2500,
			Stock:     100,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: time.Now(),
		},
		{
			Id:        "2",
			ItemName:  "item_2",
			Price:     1200,
			Stock:     200,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: time.Now(),
		},
	}
}
