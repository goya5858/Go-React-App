package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GET_all_items(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/items endpoint is hooked!")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}
