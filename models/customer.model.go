package models
// Struct untuk menyimpan data pelanggan
type Customer struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}