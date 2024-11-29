package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

//-------------EMPLOYEE---------------//
	//GET

	router.GET("/employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Get Employee",
		})
	})

	//POSE

	router.POST("/employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Past employee",
		})
	})
	//PUT

	router.PUT("/employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Put Employee",
		})
	})

	//DELETE

	router.DELETE("/employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Delete Employee",
		})
	})

//-------------CUSTOMER---------------//
	//GET

	router.GET("/customer", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Get Employee",
		})
	})

	//POSE

	router.POST("/customer", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Past employee",
		})
	})
	//PUT

	router.PUT("/customer", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Put Employee",
		})
	})

	//DELETE

	router.DELETE("/customer", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Delete Employee",
		})
	})
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}