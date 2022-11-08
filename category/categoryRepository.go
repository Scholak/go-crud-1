package category

import (
	"errors"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Index() ([]Category, error)
	Store(category *Category) error
	Show(id int) (*Category, error)
	Update(id int, name string) error
	Delete(id int) error
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
	return categories, c.db.Find(&categories).Error
}

func (c *categoryRepository) Store(category *Category) error {
	return c.db.Model(&Category{}).Create(category).Error
}

func (c *categoryRepository) Show(id int) (*Category, error) {
	category := &Category{}
	return category, c.db.Find(&category, id).Error
}

func (c *categoryRepository) Update(id int, name string) error {
	return c.db.Model(&Category{}).Where("id = ?", id).Update("name", name).Error
}

func (c *categoryRepository) Delete(id int) error {
	return c.db.Delete(&Category{}, id).Error
}