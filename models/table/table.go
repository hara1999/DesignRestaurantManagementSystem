package table

type TableStatus string

const (
	BOOKED    TableStatus = "BOOKED"
	AVAILABLE TableStatus = "AVAILABLE"
)

type Table struct {
	TableID string
	Status  TableStatus
}

func NewTable(tableID string) *Table {
	return &Table{
		TableID: tableID,
		Status:  AVAILABLE,
	}
}

func (t *Table) SetStatus(status TableStatus) {
	t.Status = status
}
