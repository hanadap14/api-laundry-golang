package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Struct untuk menyimpan data pelanggan
type Customer struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

// Struct untuk menyimpan data pesanan
type Order struct {
	ID         int     `json:"id"`
	CustomerID int     `json:"customer_id"`
	Item       string  `json:"item"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}

// Struct untuk menyimpan data tagihan
type Bill struct {
	ID      int     `json:"id"`
	OrderID int     `json:"order_id"`
	Amount  float64 `json:"amount"`
	Paid    bool    `json:"paid"`
}

// Struct untuk menyimpan data pengambilan
type Pickup struct {
	ID      int    `json:"id"`
	OrderID int    `json:"order_id"`
	Date    string `json:"date"`
}

// Struct untuk menyimpan konfigurasi database
type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
}

// Variabel untuk menyimpan instance database
var db *sql.DB

// Menghubungkan ke database
func ConnectDB(config DBConfig) {
	connectionString := config.Username + ":" + config.Password + "@tcp(" + config.Host + ":" + config.Port + ")/" + config.DBName + "?parseTime=true"
	var err error
	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
}

// Membuat tabel-tabel di database
func CreateTables() {
	query := `
		CREATE TABLE IF NOT EXISTS customers (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			address VARCHAR(255) NOT NULL
		);
		CREATE TABLE IF NOT EXISTS orders (
			id INT AUTO_INCREMENT PRIMARY KEY,
			customer_id INT NOT NULL,
			item VARCHAR(255) NOT NULL,
			quantity INT NOT NULL,
			total_price FLOAT NOT NULL,
			FOREIGN KEY (customer_id) REFERENCES customers(id)
		);
		CREATE TABLE IF NOT EXISTS bills (
			id INT AUTO_INCREMENT PRIMARY KEY,
			order_id INT NOT NULL,
			amount FLOAT NOT NULL,
			paid BOOLEAN NOT NULL,
			FOREIGN KEY (order_id) REFERENCES orders(id)
		);
		CREATE TABLE IF NOT EXISTS pickups (
			id INT AUTO_INCREMENT PRIMARY KEY,
			order_id INT NOT NULL,
			date VARCHAR(255) NOT NULL,
			FOREIGN KEY (order_id) REFERENCES orders(id)
		);
	`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

// Handler untuk mendapatkan semua data pelanggan
func GetCustomers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, address FROM customers")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	customers := []Customer{}
	for rows.Next() {
		var customer Customer
		err := rows.Scan(&customer.ID, &customer.Name, &customer.Address)
		if err != nil {
			log.Fatal(err)
		}
		customers = append(customers, customer)
	}

	json.NewEncoder(w).Encode(customers)
}

// Handler untuk membuat data pelanggan baru
func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer Customer
	json.NewDecoder(r.Body).Decode(&customer)

	insertQuery := "INSERT INTO customers (name, address) VALUES (?, ?)"
	_, err := db.Exec(insertQuery, customer.Name, customer.Address)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
}

// Handler untuk mendapatkan semua data pesanan
func GetOrders(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, customer_id, item, quantity, total_price FROM orders")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	orders := []Order{}
	for rows.Next() {
		var order Order
		err := rows.Scan(&order.ID, &order.CustomerID, &order.Item, &order.Quantity, &order.TotalPrice)
		if err != nil {
			log.Fatal(err)
		}
		orders = append(orders, order)
	}

	json.NewEncoder(w).Encode(orders)
}

// Handler untuk membuat data pesanan baru
func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order Order
	json.NewDecoder(r.Body).Decode(&order)

	insertQuery := "INSERT INTO orders (customer_id, item, quantity, total_price) VALUES (?, ?, ?, ?)"
	_, err := db.Exec(insertQuery, order.CustomerID, order.Item, order.Quantity, order.TotalPrice)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
}

// Handler untuk mendapatkan semua data tagihan
func GetBills(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, order_id, amount, paid FROM bills")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	bills := []Bill{}
	for rows.Next() {
		var bill Bill
		err := rows.Scan(&bill.ID, &bill.OrderID, &bill.Amount, &bill.Paid)
		if err != nil {
			log.Fatal(err)
		}
		bills = append(bills, bill)
	}

	json.NewEncoder(w).Encode(bills)
}

// Handler untuk membuat data tagihan baru
func CreateBill(w http.ResponseWriter, r *http.Request) {
	var bill Bill
	json.NewDecoder(r.Body).Decode(&bill)

	insertQuery := "INSERT INTO bills (order_id, amount, paid) VALUES (?, ?, ?)"
	_, err := db.Exec(insertQuery, bill.OrderID, bill.Amount, bill.Paid)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
}

// Handler untuk mendapatkan semua data pengambilan
func GetPickups(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, order_id, date FROM pickups")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	pickups := []Pickup{}
	for rows.Next() {
		var pickup Pickup
		err := rows.Scan(&pickup.ID, &pickup.OrderID, &pickup.Date)
		if err != nil {
			log.Fatal(err)
		}
		pickups = append(pickups, pickup)
	}

	json.NewEncoder(w).Encode(pickups)
}

// Handler untuk membuat data pengambilan baru
func CreatePickup(w http.ResponseWriter, r *http.Request) {
	var pickup Pickup
	json.NewDecoder(r.Body).Decode(&pickup)

	insertQuery := "INSERT INTO pickups (order_id, date) VALUES (?, ?)"
	_, err := db.Exec(insertQuery, pickup.OrderID, pickup.Date)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
}

// Handler untuk menghapus data pengambilan
func DeletePickup(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pickupID := params["id"]

	deleteQuery := "DELETE FROM pickups WHERE id = ?"
	_, err := db.Exec(deleteQuery, pickupID)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
}

func main() {
	config := DBConfig{
		Username: "your_username",
		Password: "your_password",
		Host:     "localhost",
		Port:     "3306",
		DBName:   "laundry_db",
	}

	ConnectDB(config)
	CreateTables()

	router := mux.NewRouter()

	// Customer endpoints
	router.HandleFunc("/customers", GetCustomers).Methods("GET")
	router.HandleFunc("/customers", CreateCustomer).Methods("POST")

	// Order endpoints
	router.HandleFunc("/orders", GetOrders).Methods("GET")
	router.HandleFunc("/orders", CreateOrder).Methods("POST")

	// Bill endpoints
	router.HandleFunc("/bills", GetBills).Methods("GET")
	router.HandleFunc("/bills", CreateBill).Methods("POST")

	// Pickup endpoints
	router.HandleFunc("/pickups", GetPickups).Methods("GET")
	router.HandleFunc("/pickups", CreatePickup).Methods("POST")
	router.HandleFunc("/pickups/{id}", DeletePickup).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
