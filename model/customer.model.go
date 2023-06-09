package models
// Struct untuk menyimpan data pelanggan
type Customer struct {
	ID      int    `json:"id" gorm:"primary_key"`
	Name    string `json:"name"`
	Address string `json:"address"`
}