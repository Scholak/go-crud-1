package main

import (
	"log"
	"net/http"

	"github.com/Scholak/go-crud/category"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.StrictSlash(true)

	c := r.PathPrefix("/api/categories").Subrouter()

	c.HandleFunc("/", category.GetCategories).Methods("GET")
	c.HandleFunc("/store", category.StoreCategory).Methods("POST")
	c.HandleFunc("/{id}", category.GetCategoryById).Methods("GET")
	c.HandleFunc("/{id}", category.UpdateCategory).Methods("PUT")
	c.HandleFunc("/{id}", category.DeleteCategory).Methods("DELETE")
	
	log.Fatal(http.ListenAndServe(":8000", r))
}