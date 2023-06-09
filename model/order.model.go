package models

// Struct untuk menyimpan data pesanan
type Order struct {
	ID         int     `json:"id"`
	CustomerID int     `json:"customer_id"`
	Item       string  `json:"item"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}