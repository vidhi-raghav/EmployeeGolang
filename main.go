package main

import (
	"employee/controller"
	"employee/employee"
	"employee/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	employeeService    service.EmployeeService       = service.New()
	employeeController controller.EmployeeController = controller.New(employeeService)
)

func main() {
	emp := employee.Employee{123, "VIDHI", "9876543210"}
	fmt.Print(emp)

	router := gin.Default()
	// router.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	router.GET("/empoyee", func(ctx *gin.Context) {
		ctx.JSON(200, employeeController.FindAll())
	})

	router.POST("/empoyee", func(ctx *gin.Context) {
		ctx.JSON(200, employeeController.Save(ctx))
	})

	router.Run()
}
