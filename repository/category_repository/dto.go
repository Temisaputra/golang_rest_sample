package category_repository

import "github.com/Temisaputra/warOnk/core/entity"

type Category struct {
	CategoryID   int    `json:"category_id" gorm:"column:id"`
	CategoryName string `json:"category_name"`
}

func (c Category) TableName() string {
	return "category_product"
}

func (c *Category) ToEntity() *entity.Category {
	return &entity.Category{
		CategoryID:   c.CategoryID,
		CategoryName: c.CategoryName,
	}
}
