package category

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Index() ([]Category, error)
	Store(category *Category) error
	Show(id int) (*Category, error)
	Update(category *Category) error
	Delete(category *Category) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) (CategoryRepository, error) {
	if db == nil {
		return nil, errors.New("db is nil #1076")
	}

	if err := db.AutoMigrate(&Category{}); err != nil {
		return nil, err
	}

	return &categoryRepository{db: db}, nil
}

func (c *categoryRepository) Index() ([]Category, error) {
	categories := []Category{}
	fmt.Println("a")
	return categories, c.db.Find(&categories).Error
}

func (c *categoryRepository) Store(category *Category) error {
	return c.db.Model(&Category{}).Create(category).Error
}

func (c *categoryRepository) Show(id int) (*Category, error) {
	category := &Category{}
	return category, c.db.Find(&category, id).Error
}

func (c *categoryRepository) Update(category *Category) error {
	return c.db.Save(&category).Error
}

func (c *categoryRepository) Delete(category *Category) error {
	return c.db.Delete(&category).Error
}