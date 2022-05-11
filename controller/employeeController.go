package controller

import (
	"employee/employee"
	"employee/service"

	"github.com/gin-gonic/gin"
)

type EmployeeController interface {
	FindAll() []employee.Employee
	Save(ctx *gin.Context) employee.Employee
}

type controller struct {
	service service.EmployeeService
}

func New(service service.EmployeeService) controller {
	return controller{
		service: service,
	}
}

func (c *controller) FindAll() []employee.Employee {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) employee.Employee {
	var employee employee.Employee
	ctx.Bind(&employee)
	c.service.Save(employee)
	return employee
}
