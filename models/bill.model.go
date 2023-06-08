package models

// Struct untuk menyimpan data tagihan
type Bill struct {
	Id      int     `json:"id" gorm:"primary_key"`
	Order_id int     `json:"order_id" gorm:"foreignkey:Order_id"`
	Amount  float64 `json:"amount"`
	Paid    bool    `json:"paid "`
}