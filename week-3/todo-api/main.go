package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var todos = map[int]todo{}

func main() {
	r := gin.Default()

	r.POST("/api/todos", createToDoHandler)
	r.GET("/api/todos", getTodosHandler)
	r.GET("/api/todos/:id", getTodoHandler)
	r.PUT("/api/todos/:id", updateTodoHandler)
	r.DELETE("/api/todos/:id", deleteTodoHandler)

	r.Run(":1234")
}

func createToDoHandler(c *gin.Context) {
	var todo todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id format",
		})
		return
	}
	// generate id
	id := len(todos)
	id++

	// assign to map
	todo.ID = id
	todos[id] = todo

	c.JSON(http.StatusCreated, todo)
}

func getTodosHandler(c *gin.Context) {
	t := []todo{}

	for _, v := range todos {
		t = append(t, v)
	}

	c.JSON(http.StatusOK, t)
}

func getTodoHandler(c *gin.Context) {
	id := c.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id format",
		})
		return
	}

	// check if exist
	_, exist := todos[i]
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "todo not found",
		})
		return
	}

	todo := todos[i]

	c.JSON(http.StatusOK, todo)
}

func updateTodoHandler(c *gin.Context) {
	id := c.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id format",
		})
		return
	}

	// check if exist
	_, exist := todos[i]
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "todo not found",
		})
		return
	}

	var t todo
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "JSON parse error",
		})
		return
	}

	// update value
	todo := todos[i]
	todo.Title = t.Title
	todo.Status = t.Status
	todos[i] = todo

	c.JSON(http.StatusOK, todo)

}

func deleteTodoHandler(c *gin.Context) {
	id := c.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// check if exist
	_, exist := todos[i]
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "todo not found",
		})
		return
	}

	// delete it!
	delete(todos, i)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
