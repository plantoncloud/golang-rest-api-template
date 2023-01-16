package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// GET all routes
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"GET":    []string{"/todos", "todos/:id"},
			"POST":   []string{"/todos"},
			"PUT":    []string{"todos/:id"},
			"DELETE": []string{"todos/:id"},
		})
	})

	// GET a single todo
	r.GET("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{"id": id})
	})

	// GET all todos
	r.GET("/todos", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"todos": []string{"todo 1", "todo 2"}})
	})

	// POST a new todo
	r.POST("/todos", func(c *gin.Context) {
		todo := c.PostForm("todos")
		c.JSON(http.StatusOK, gin.H{"todo": todo})
	})

	// PUT an updated todo
	r.PUT("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")
		todo := c.PostForm("todo")
		c.JSON(http.StatusOK, gin.H{"id": id, "todo": todo})
	})

	// DELETE a todo
	r.DELETE("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{"id": id})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
