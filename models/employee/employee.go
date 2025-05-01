package employee

type Designation string

const (
	CHEF    Designation = "CHEF"
	WAITER  Designation = "WAITER"
	MANAGER Designation = "MANAGER"
)

type WorkShift string

const (
	NIGHT WorkShift = "NIGHT"
	DAY   WorkShift = "DAY"
)

type Employee struct {
	Name        string
	Designation Designation
	WorkShift   WorkShift
}

func NewEmployee(name string, designation Designation) *Employee {
	return &Employee{
		Name:        name,
		Designation: designation,
	}
}
