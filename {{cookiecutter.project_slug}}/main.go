package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// GET a single task
	r.GET("/task/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{"id": id})
	})

	// GET all tasks
	r.GET("/tasks", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"tasks": []string{"task 1", "task 2"}})
	})

	// POST a new task
	r.POST("/task", func(c *gin.Context) {
		task := c.PostForm("task")
		c.JSON(http.StatusOK, gin.H{"task": task})
	})

	// PUT an updated task
	r.PUT("/task/:id", func(c *gin.Context) {
		id := c.Param("id")
		task := c.PostForm("task")
		c.JSON(http.StatusOK, gin.H{"id": id, "task": task})
	})

	// DELETE a task
	r.DELETE("/task/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{"id": id})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
