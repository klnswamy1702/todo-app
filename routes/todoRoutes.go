package routes

import (
    "go-todo-app/controllers"
    "github.com/gorilla/mux"
)

func TodoRoutes(router *mux.Router) {
    router.HandleFunc("/todos", controllers.CreateTodoEndpoint).Methods("POST")
    router.HandleFunc("/todos", controllers.GetTodosEndpoint).Methods("GET")
    router.HandleFunc("/todos/{id}", controllers.UpdateTodoEndpoint).Methods("PUT")
    router.HandleFunc("/todos/{id}", controllers.DeleteTodoEndpoint).Methods("DELETE")
}

