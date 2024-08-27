package main

import (
    "github.com/gin-gonic/gin"
    "github.com/klnswamy1702/todo-app/backend/config"
    "github.com/klnswamy1702/todo-app/backend/controllers"
    "github.com/klnswamy1702/todo-app/backend/repositories"
    "github.com/klnswamy1702/todo-app/backend/routes"
    "github.com/klnswamy1702/todo-app/backend/services"
)

func main() {
    config.ConnectDB()

    repo := repositories.NewTodoRepository()
    service := services.NewTodoService(repo)
    controller := controllers.NewTodoController(service)

    router := gin.Default()
    routes.RegisterTodoRoutes(router, controller)

    router.Run(":8080")
}
