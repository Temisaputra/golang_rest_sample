package unit_repository

import "github.com/Temisaputra/warOnk/core/entity"

type ProductUnit struct {
	ID       int    `json:"id" gorm:"column:id"` // gorm:"column:unit_id" is used to map the column name in the database
	UnitName string `json:"unit_name" gorm:"column:unit_name"`
}

func (u ProductUnit) TableName() string {
	return "product_unit"
}

func (u *ProductUnit) ToEntity() *entity.ProductUnit {
	return &entity.ProductUnit{
		UnitID:   u.ID,
		UnitName: u.UnitName,
	}
}
