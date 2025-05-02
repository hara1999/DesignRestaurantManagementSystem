package item

type AddItemRequest struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}
