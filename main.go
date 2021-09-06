package main

import (
	"github.com/gorilla/mux"
	"github.com/swapna-pg/golang/data-collection-tree/handler"
	"net/http"
)

func main() {
	dataCollector := handler.NewDataCollector()

	router := mux.NewRouter()
	router.HandleFunc("/v1/insert", dataCollector.Insert).Methods("POST")
	router.HandleFunc("/v1/query", dataCollector.Query).Methods("GET")
	http.ListenAndServe(":8080", router)
}
