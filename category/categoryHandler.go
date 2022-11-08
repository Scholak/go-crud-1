package category

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Scholak/go-crud/config"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error
var categoryRepo CategoryRepository

func init() {
	db, err = config.Connection()
	if err != nil {
		fmt.Println("DB connection error:", err)
		return
	}
	
	categoryRepo, err = NewCategoryRepository(db)
	if err != nil {
		fmt.Println("Category repo error:", err)
		return
	}
}

func GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := categoryRepo.Index()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	
	response := map[string]interface{}{
		"categories": categories,
	}
	json.NewEncoder(w).Encode(response)
}

func StoreCategory(w http.ResponseWriter, r *http.Request) {
	var newCategory *Category
	json.NewDecoder(r.Body).Decode(&newCategory)
	
	if newCategory.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Category name cannot be empty"))
		return
	}
	
	if err := categoryRepo.Store(newCategory); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	
	response := map[string]interface{}{
		"message": "Category created successfully!",
		"category": newCategory,
	}
	json.NewEncoder(w).Encode(response)
}

func GetCategoryById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	
	if !categoryRepo.CheckExistance(id) {
		response := map[string]interface{}{
			"message": "Category not found!",
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}
	
	category, err := categoryRepo.Show(id)
	
	response := map[string]interface{}{
		"category": category,
	}
	json.NewEncoder(w).Encode(response)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	var updatedCategory *Category
	
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if !categoryRepo.CheckExistance(id) {
		response := map[string]interface{}{
			"message": "Category not found!",
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}
	
	json.NewDecoder(r.Body).Decode(&updatedCategory)
	
	if updatedCategory.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Category name cannot be empty"))
		return
	}
	
	err = categoryRepo.Update(updatedCategory, id)
	
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	
	response := map[string]interface{}{
		"message": "Category updated successfully!",
		"category": updatedCategory,
	}
	json.NewEncoder(w).Encode(response)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if !categoryRepo.CheckExistance(id) {
		response := map[string]interface{}{
			"message": "Category not found!",
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}
	
	if err := categoryRepo.Delete(id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	
	response := map[string]interface{}{
		"message": "Category deleted successfully!",
	}
	json.NewEncoder(w).Encode(response)
}