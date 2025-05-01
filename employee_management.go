package main

import "errors"

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

func NewStaff(name string, designation Designation) *Employee {
	return &Employee{
		Name:        name,
		Designation: designation,
	}
}

type EmployeeManagementService struct {
	employees []*Employee
}

func (em *EmployeeManagementService) EnrollEmployee(emp *Employee) {
	em.employees = append(em.employees, emp)
}

func (em *EmployeeManagementService) GetEmployeeByName(name string) (*Employee, error) {
	for _, emp := range em.employees {
		if emp.Name == name {
			return emp, nil
		}
	}
	return nil, errors.New("no employee found with provided name")
}

func (em *EmployeeManagementService) GetEmployeesByDesignation(designation Designation) ([]*Employee, error) {
	var emps []*Employee
	for _, emp := range em.employees {
		if emp.Designation == designation {
			emps = append(emps, emp)
		}
	}

	if len(emps) > 0 {
		return emps, nil
	}
	return nil, errors.New("no employee found with provided designation")
}

func (em *EmployeeManagementService) AssignShift(employeeName string, shift WorkShift) error {
	for _, employee := range em.employees {
		if employee.Name == employeeName {
			employee.WorkShift = shift
			return nil
		}
	}
	return errors.New("employee or shift doesn't exist")
}
