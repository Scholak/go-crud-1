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
	c.HandleFunc("/create", category.StoreCategory).Methods("POST")
	c.HandleFunc("/{id}", category.GetCategoryById).Methods("GET")
	c.HandleFunc("/{id}/edit", category.UpdateCategory).Methods("PUT")
	c.HandleFunc("/{id}/delete", category.DeleteCategory).Methods("DELETE")

	// category.GetCategories()
	// c, err := category.GetCategoryById(1)
	// if err != nil {
	// 	panic(err)
	// }
	
	// c.Name = "updated category"
	
	// category.UpdateCategory(c)
	// fmt.Println(category.GetCategoryById(1))
	
	// category.DeleteCategory(c)
	// fmt.Println(category.GetCategoryById(1))
	
	log.Fatal(http.ListenAndServe(":8000", r))
}