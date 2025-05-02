package item

type ItemStatus string

const (
	ItemStatusGood    ItemStatus = "good"
	ItemStatusExpired ItemStatus = "expired"
)

type Item struct {
	ItemName string
	Quantity int
	Status   ItemStatus
}

func NewItem(itemName string, quantity int) *Item {
	return &Item{
		ItemName: itemName,
		Quantity: quantity,
		Status:   ItemStatusGood,
	}
}
