package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// GET a single todo
	r.GET("/todo/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{"id": id})
	})

	// GET all todos
	r.GET("/todos", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"todos": []string{"todo 1", "todo 2"}})
	})

	// POST a new todo
	r.POST("/todo", func(c *gin.Context) {
		todo := c.PostForm("todo")
		c.JSON(http.StatusOK, gin.H{"todo": todo})
	})

	// PUT an updated todo
	r.PUT("/todo/:id", func(c *gin.Context) {
		id := c.Param("id")
		todo := c.PostForm("todo")
		c.JSON(http.StatusOK, gin.H{"id": id, "todo": todo})
	})

	// DELETE a todo
	r.DELETE("/todo/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{"id": id})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
