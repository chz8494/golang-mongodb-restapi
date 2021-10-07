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

	apiRouter := mux.NewRouter().StrictSlash(true)
	apiRouter.HandleFunc("/api/articles", controllers.CreateArticle).Methods("POST")
	apiRouter.HandleFunc("/api/articles/", controllers.GetArticles).Methods("GET")
	apiRouter.HandleFunc("/api/articles/{id}", controllers.GetArticle).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), apiRouter))
}

// entry point
func main() {
	// connect db

	handleRequests()
}
