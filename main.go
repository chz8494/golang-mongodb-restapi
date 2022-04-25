package main

import (
	"github.com/chz8494/golang-mongodb-restapi/controllers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
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
	apiRouter.HandleFunc("/api/coins/{coin}", controllers.GetCoin).Methods("GET")
	apiRouter.HandleFunc("/api/coins/{coin}/{timestamp}", controllers.GetCoin_Timestamp).Methods("GET")
	corsObj := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(corsObj)(apiRouter))
}

// entry point
func main() {
	// connect db

	handleRequests()
}
