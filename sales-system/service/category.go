package service

import (
	"gorm.io/gorm"
	"microservice/sales-system/model"
)

type Category struct {
	DB *gorm.DB
}

func (ca *Category) Add(name string) (int32, error) {
	category := model.Category{
		CateName: name,
	}
	result := ca.DB.Create(&category)

	if result.Error != nil {
		return 0, result.Error
	}

	return category.Id, nil
}
