package models

// Struct untuk menyimpan data tagihan
type Bill struct {
	ID      int     `json:"id"`
	OrderID int     `json:"order_id"`
	Amount  float64 `json:"amount"`
	Paid    bool    `json:"paid"`
}