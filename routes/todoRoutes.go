package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/klnswamy1702/todo-app/backend/controllers"
)

func RegisterTodoRoutes(router *gin.Engine, controller *controllers.TodoController) {
    router.GET("/todos", controller.GetTodos)
    router.GET("/todos/:id", controller.GetTodoByID)
    router.POST("/todos", controller.CreateTodo)
    router.PUT("/todos/:id", controller.UpdateTodo)
    router.DELETE("/todos/:id", controller.DeleteTodo)
}
