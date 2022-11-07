package category

import (
	"fmt"

	"github.com/Scholak/go-crud/config"
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

func GetCategories() {
	categories, err := categoryRepo.Index()

	if err != nil {
		fmt.Println(err)
        return
	}

	fmt.Println(categories)
}

func StoreCategory() {
	c := newCategory("category one")

	if err := categoryRepo.Store(c); err != nil {
		fmt.Println("Failed to insert record!", err)
		return
	}
}

func GetCategoryById(id int) (*Category, error) {
	category, err := categoryRepo.Show(id)
	
	if err != nil {
		fmt.Println(err)
        return nil, err
	}

	return category, nil
}

func UpdateCategory(category *Category) {
	err := categoryRepo.Update(category)

	if err != nil {
		fmt.Println(err)
        return
	}
}

func DeleteCategory(category *Category) {
	err := categoryRepo.Delete(category)

	if err != nil {
		fmt.Println(err)
        return
	}
}