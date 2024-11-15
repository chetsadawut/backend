package main

import (
	"net/http"
  
	"github.com/gin-gonic/gin"
  )

  func main() {
	r := gin.Default()

	r.GET("/employee check", func(c *gin.Context) {
	  c.JSON(http.StatusOK, gin.H{
		"message": "TEST EMPLOYEE CHECK",
	  })
	})

	//GET METHOD
	r.GET("/employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
		  "message": "GET EMPLOYEE METHOD",
		})
	  })

	//POST METHOD
	r.POST("/employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
		  "message": "POST EMPLOYEE METHOD",
		})
	  })

	  //PUT METHOD
	  r.PUT("/employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
		  "message": "PUT EMPLOYEE METHOD",
		})
	  })

	  //DELETE METHOD
	  r.DELETE("/employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
		  "message": "DELETE EMPLOYEE METHOD",
		})
	  })

///////////////////CUSTOMER CHECK///////////////////

	r.GET("/customer check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
		  "message": "TEST EMPLOYEE CHECK",
		})
	  })
	//GET METHOD
	r.GET("/customer", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
		  "message": "GET CUSTOMER METHOD",
		})
	  })

	//POST METHOD
	r.POST("/customer", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
		  "message": "POST CUSTOMER METHOD",
		})
	  })

	  //PUT METHOD
	  r.PUT("/customer", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
		  "message": "PUT CUSTOMER METHOD",
		})
	  })

	  //DELETE METHOD
	  r.DELETE("/customer", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
		  "message": "DELETE CUSTOMER METHOD",
		})
	  })
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
  }