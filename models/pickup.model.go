package models

// Struct untuk menyimpan data pengambilan
type Pickup struct {
	ID      int    `json:"id"`
	OrderID int    `json:"order_id"`
	Date    string `json:"date"`
}