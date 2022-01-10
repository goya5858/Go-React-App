package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func StartWebServer() error {
	fmt.Println("Rest API with Mux Routers")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", rootPage)
	router.HandleFunc("/items", GET_all_items).Methods("GET")
	router.HandleFunc("/item/{id}", GET_one_item).Methods("GET")

	router.HandleFunc("/item", POST_item).Methods("POST")
	router.HandleFunc("/item/{id}", DELETE_item).Methods("DELETE")
	router.HandleFunc("/item/{id}", PUT_item).Methods("PUT")

	return http.ListenAndServe(fmt.Sprintf(":%d", 8080), router)
}

func rootPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Go Api Server")
	fmt.Println("Root endpoint is hooked!")
}

type ItemParams struct {
	Id       string `json:"id"`
	ItemName string `json:"item_name,omitempty"`
	Price    int    `json:"price,omitempty"`
	Stock    int    `json:"stock,omitempty"`
}

var items []*ItemParams

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
