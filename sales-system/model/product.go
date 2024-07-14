package model

type Product struct {
	Id               int32   `json:"id" gorm:"primaryKey"`
	ProdName         string  `json:"prod_name" gorm:"type:varchar(50);not null;unique;comment:商品名称"`
	BaseUnit         string  `json:"base_unit" gorm:"type:varchar(10);not null;comment:基本单位"`
	UnitPrice        float32 `json:"unit_price" gorm:"type:int;not null;comment:单价"`
	PurchaseQuantity int32   `json:"purchase_quantity" gorm:"type:int;not null;comment:采购数量"`
	TotalPrice       float32 `json:"total_price" gorm:"type:int;not null;comment:总价"`
	GormModel
}
