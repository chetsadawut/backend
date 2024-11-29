package main

import (
	EmployeeController "backend/api/controller/employee"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//GET
	router.GET("/employee", EmployeeController.GetEmployee)
	//POST
	router.POST("/employee", EmployeeController.PostEmployee)
	//PUT
	router.PUT("/employee", EmployeeController.PutEmployee)
	//DELETE
	router.DELETE("/employee", EmployeeController.DeleteEmployee)
	//GET By ID
	router.GET("/employee/:id", EmployeeController.GetEmployeeByID)
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
