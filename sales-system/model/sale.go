package model

type Sale struct {
	ID            int32   `json:"id" gorm:"primaryKey"`
	ProdName      string  `json:"prod_name" gorm:"type:varchar(50);not null;comment:商品名称"`
	BaseUnit      string  `json:"base_unit" gorm:"type:varchar(20);not null;comment:基本单位"`
	SellingVolume float32 `json:"selling_volume" gorm:"type:float;not null;comment:销售数量"`
	TotalAmount   float32 `json:"total_amount" gorm:"type:float;not null;comment:总金额"`
	CreateDate    int64   `json:"create_date" gorm:"type:int;not null;comment:销售日期"`
	Remark        string  `json:"remark" gorm:"type:varchar(255);comment:备注"`
	GormModel
}
