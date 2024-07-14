package model

type Category struct {
	ID       int32  `json:"id" gorm:"primaryKey"`
	CateName string `json:"cate_name" gorm:"type:varchar(50);not null;unique;comment:商品种类"`
	GormModel
}
