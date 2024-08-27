package controllers

import (
    "github.com/gin-gonic/gin"
    "github.com/klnswamy1702/todo-app/backend/models"
    "github.com/klnswamy1702/todo-app/backend/services"
    "net/http"
)

type TodoController struct {
    service services.TodoService
}

func NewTodoController(service services.TodoService) *TodoController {
    return &TodoController{service: service}
}

func (c *TodoController) GetTodos(ctx *gin.Context) {
    todos, err := c.service.GetAllTodos()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, todos)
}

func (c *TodoController) GetTodoByID(ctx *gin.Context) {
    id := ctx.Param("id")
    todo, err := c.service.GetTodoByID(id)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
        return
    }
    ctx.JSON(http.StatusOK, todo)
}

func (c *TodoController) CreateTodo(ctx *gin.Context) {
    var todo models.Todo
    if err := ctx.ShouldBindJSON(&todo); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := c.service.CreateTodo(todo); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusCreated, todo)
}

func (c *TodoController) UpdateTodo(ctx *gin.Context) {
    id := ctx.Param("id")
    var todo models.Todo
    if err := ctx.ShouldBindJSON(&todo); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := c.service.UpdateTodo(id, todo); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, todo)
}

func (c *TodoController) DeleteTodo(ctx *gin.Context) {
    id := ctx.Param("id")
    if err := c.service.DeleteTodo(id); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
