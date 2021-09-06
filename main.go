package main

import (
	"github.com/gorilla/mux"
	"github.com/swapna-pg/golang/data-collection-tree/handler"
	"net/http"
)

func main() {
	dataCollector := handler.NewDataCollector()
	dataCollector.StartConsuming()

	router := mux.NewRouter()
	api := router.PathPrefix("/").Subrouter()
	v1API:= api.PathPrefix("/v1").Subrouter()
	v1API.HandleFunc("/insert", dataCollector.Insert).Methods("POST")
	v1API.HandleFunc("/query", dataCollector.Query).Methods("GET")
	http.ListenAndServe(":8080", router)
}
