package service

import (
	"errors"

	employeeModel "DesignRestaurantManagementSystem/models/employee"
)

type EmployeeManagementService struct {
	employees []*employeeModel.Employee
}

func NewEmployeeManagementService() *EmployeeManagementService {
	return &EmployeeManagementService{
		employees: make([]*employeeModel.Employee, 0),
	}
}

func (em *EmployeeManagementService) EnrollEmployee(emp *employeeModel.Employee) {
	em.employees = append(em.employees, emp)
}

func (em *EmployeeManagementService) GetEmployeeByName(name string) (*employeeModel.Employee, error) {
	for _, emp := range em.employees {
		if emp.Name == name {
			return emp, nil
		}
	}
	return nil, errors.New("no employee found with provided name")
}

func (em *EmployeeManagementService) GetEmployeesByDesignation(designation employeeModel.Designation) ([]*employeeModel.Employee, error) {
	var emps []*employeeModel.Employee
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

func (em *EmployeeManagementService) AssignShift(employeeName string, shift employeeModel.WorkShift) error {
	for _, employee := range em.employees {
		if employee.Name == employeeName {
			employee.WorkShift = shift
			return nil
		}
	}
	return errors.New("employee or shift doesn't exist")
}
