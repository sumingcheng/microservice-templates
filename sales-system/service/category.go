package service

import (
	"gorm.io/gorm"
	"microservice/sales-system/model"
	"microservice/sales-system/utils"
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

type CateList struct {
	Count int64            `json:"count"`
	Data  []model.Category `json:"data"`
}

func (ca *Category) GetList(pageSize, pageNumber int) (*CateList, error) {
	var cateList []model.Category
	result := ca.DB.Find(&cateList)

	if result.Error != nil {
		return nil, result.Error
	}

	count := result.RowsAffected
	ca.DB.Scopes(utils.Paginate(int32(pageNumber), int32(pageSize))).Find(&cateList)

	return &CateList{
		Count: count,
		Data:  cateList,
	}, nil
}

func (ca *Category) Exists(name string) (bool, error) {
	var count int64
	result := ca.DB.Model(&model.Category{}).Where("cate_name = ?", name).Count(&count)

	if result.Error != nil {
		return false, result.Error
	}

	return count > 0, nil
}

func (ca *Category) GetOne(id int) (*model.Category, error) {
	var category model.Category
	result := ca.DB.Where("id = ?", id).First(&category)

	if result.Error != nil {
		return nil, result.Error
	}

	return &category, nil
}
