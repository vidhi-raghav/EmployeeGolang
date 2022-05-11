package service

import "employee/employee"

type EmployeeService interface {
	Save(employee.Employee) employee.Employee
	FindAll() []employee.Employee
}

type employeeService struct {
	employees []employee.Employee
}

func New() EmployeeService {
	return &employeeService{}
}

func (service *employeeService) Save(employee employee.Employee) employee.Employee {
	service.employees = append(service.employees, employee)
	return employee
}

func (service *employeeService) FindAll() []employee.Employee {
	return service.employees
}
