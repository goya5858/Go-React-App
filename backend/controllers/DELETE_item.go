package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func DELETE_item(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, item := range items {
		if item.Id == id {
			// indexがidの部分を飛ばす
			items = append(items[:index], items[index+1:]...)
		}
	}
}
