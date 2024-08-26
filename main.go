package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "go-todo-app/routes"
    "go-todo-app/controllers"
    "github.com/gorilla/mux"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }
    err = client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Connected to MongoDB!")
    return client
}

func main() {
    client := ConnectDB()
    collection := client.Database("tododb").Collection("todos")
    controllers.SetCollection(collection)

    router := mux.NewRouter()
    routes.TodoRoutes(router)

    fmt.Println("Starting server on the port 8000...")
    log.Fatal(http.ListenAndServe(":8000", router))
}

