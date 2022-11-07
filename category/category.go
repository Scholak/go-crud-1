package category

type Category struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"index"`
	Hit int `json:"hit"`
}

func newCategory(name string) *Category {
	return &Category{
		Name: name,
		Hit: 0,
	}
}