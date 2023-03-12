package requests

type AddProductToCartRequest struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}
