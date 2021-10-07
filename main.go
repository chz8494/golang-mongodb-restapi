package main

import (
	"log"
	"net/http"
  "os"
	"github.com/chz8494/golang-mongodb-restapi/controllers"
	"github.com/gorilla/mux"
)

//handle API Requests
func handleRequests() {
	port := os.Getenv("PORT")
	if port == "" {
			 port = "8081"
	}
	apiRouter := mux.NewRouter().StrictSlash(true)
	apiRouter.HandleFunc("/api/coins", controllers.CreateCoin).Methods("POST")
	apiRouter.HandleFunc("/api/coins/", controllers.GetCoins).Methods("GET")
	apiRouter.HandleFunc("/api/coins/{id}", controllers.GetCoin).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+port, apiRouter))
}

// entry point
func main() {
	// connect db

	handleRequests()
}
