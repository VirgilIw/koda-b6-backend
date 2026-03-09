package dto

type Product struct {
	Id                int
	Name              string
	Description       string
	Price             float64
	IsBuy1Get1        bool
	IsFlashSale       bool
	IsBirthdayPackage bool
}
type ProductsResponse struct {
	Success bool
	Message string
	Error   string
	Result  []Product
}
