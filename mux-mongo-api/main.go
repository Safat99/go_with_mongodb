package main

import (
	"log"
	"mux-mongo-api/configs"
	"mux-mongo-api/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Application Entry Point
	// router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	// 	rw.Header().Set("Content-Type", "application/json")
	// 	json.NewEncoder(rw).Encode(map[string]string{"data": "hello from mux and mongoDB"})
	// }).Methods("GET")

	//run Database
	configs.ConnectDB()

	// routes
	routes.UserRoute(router)

	log.Fatal(http.ListenAndServe(":6000", router))

}
