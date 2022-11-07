package main

import (
	"fmt"

	"github.com/Scholak/go-crud/category"
)

func main() {
	category.GetCategories()
	c, err := category.GetCategoryById(1)
	if err != nil {
		panic(err)
	}
	
	c.Name = "updated category"
	
	category.UpdateCategory(c)
	fmt.Println(category.GetCategoryById(1))
	
	category.DeleteCategory(c)
	fmt.Println(category.GetCategoryById(1))
	
	
}