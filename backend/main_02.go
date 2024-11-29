package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//Employee API method
	router.GET("/employee1", getEmployee)
	router.POST("/employee2", postEmployee)
	router.PUT("/employee3", putEmployee)
	router.DELETE("/employee4", deleteEmployee)


	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	}

func getEmployee(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{
  "message": "Employee GET Method!",
  })
}

func postEmployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
	"message": "Employee POST Method!",
	})
  }

func putEmployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
	"message": "Employee PUT Method!",
	})
  }

func deleteEmployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
	"message": "Employee DELETE Method!",
	})
  }