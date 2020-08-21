package main

import (
	"controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//handle API Requests
func handleRequests() {

	apiRouter := mux.NewRouter().StrictSlash(true)
	apiRouter.HandleFunc("/api/articles", controllers.CreateArticle).Methods("POST")
	apiRouter.HandleFunc("/api/articles/", controllers.GetArticles).Methods("GET")
	apiRouter.HandleFunc("/api/articles/{id}", controllers.GetArticle).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", apiRouter))
}

// entry point
func main() {

	handleRequests()
}
