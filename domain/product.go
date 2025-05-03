package domain

type Product struct {
	ID          int
	Name        string
	Brand       string
	Category    string
	Price       float64
	Quantity    int
	ExpiryDate  string
	SKU         string
}